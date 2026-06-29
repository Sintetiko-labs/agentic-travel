#!/usr/bin/env python3
"""Load/save scripts/groups.json with parent slug and brand_parents maps."""

from __future__ import annotations

import json
from pathlib import Path
from typing import Any

ROOT = Path(__file__).resolve().parent.parent
GROUPS_PATH = ROOT / "scripts" / "groups.json"

# CLI slug -> API parent slug when the booking API is owned by another CLI.
# Push updates here first so parallel parent-API agents can coordinate.
PARENT_OVERRIDES: dict[str, str] = {
    # Hotels — shared corporate booking APIs
    "designhotels": "marriott",
    "leonardo": "radisson",
    "mamashelter": "accor",
    "25hours": "accor",
    # Airlines — tour/charter groups sharing parent search
    "iberojet": "tui",
}


def load_groups_doc(path: Path | None = None) -> dict[str, Any]:
    path = path or GROUPS_PATH
    raw = json.loads(path.read_text(encoding="utf-8"))
    if isinstance(raw, list):
        return {"version": 1, "parent_overrides": dict(PARENT_OVERRIDES), "groups": raw}
    if "parent_overrides" not in raw:
        raw["parent_overrides"] = dict(PARENT_OVERRIDES)
    return raw


def groups_list(doc: dict[str, Any]) -> list[dict[str, Any]]:
    return doc["groups"] if "groups" in doc else doc


def parent_for_slug(slug: str, doc: dict[str, Any]) -> str:
    overrides = {**PARENT_OVERRIDES, **doc.get("parent_overrides", {})}
    return overrides.get(slug, slug)


def enrich_groups(doc: dict[str, Any]) -> dict[str, Any]:
    overrides = {**PARENT_OVERRIDES, **doc.get("parent_overrides", {})}
    groups = groups_list(doc)
    brand_parents: dict[str, str] = {}

    for g in groups:
        slug = g["slug"]
        parent = overrides.get(slug, g.get("parent", slug))
        g["parent"] = parent
        for brand in g["brands"]:
            brand_parents[brand] = parent

    doc["version"] = doc.get("version", 1)
    doc["parent_overrides"] = overrides
    doc["groups"] = groups
    doc["brand_parents"] = brand_parents
    return doc


def save_groups_doc(doc: dict[str, Any], path: Path | None = None) -> None:
    path = path or GROUPS_PATH
    enriched = enrich_groups(doc)
    path.write_text(json.dumps(enriched, indent=2, ensure_ascii=False) + "\n", encoding="utf-8")


def slug_to_group(doc: dict[str, Any]) -> dict[str, dict[str, Any]]:
    return {g["slug"]: g for g in groups_list(doc)}


def default_brand(group: dict[str, Any]) -> str:
    brands = group.get("brands") or []
    return brands[0] if brands else group["name"]
