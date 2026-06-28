#!/usr/bin/env python3
"""Wire travelkit network doctor + --no-proxy into all *-cli/cmd/*/main.go."""

from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
NETWORK_BODY = (ROOT / "scripts/templates/network.go").read_text()


def patch_main(path: Path, slug: str) -> bool:
    text = path.read_text()
    if "network.PreprocessArgs" in text and 'case "network":' in text:
        return False

    if '"github.com/fbelchi/travelkit/network"' not in text:
        text = text.replace(
            'import (\n\t"fmt"\n\t"os"\n)',
            'import (\n\t"fmt"\n\t"os"\n\n\t"github.com/fbelchi/travelkit/network"\n)',
        )

    if "network.PreprocessArgs" not in text:
        text = text.replace(
            "func main() {\n\tif len(os.Args) < 2 {",
            "func main() {\n\tos.Args = network.PreprocessArgs(os.Args)\n\tif len(os.Args) < 2 {",
        )

    if 'case "network":' not in text:
        for anchor in ('\tcase "brands":', '\tcase "version"'):
            if anchor in text:
                text = text.replace(
                    anchor,
                    '\tcase "network":\n\t\terr = cmdNetwork(os.Args[2:])\n' + anchor,
                    1,
                )
                break

    usage_line = f"  {slug} network doctor [--json]"
    if "network doctor" not in text:
        for old in (
            f"  {slug} version | help\n",
            f"  {slug} brands\n  {slug} version | help\n",
        ):
            if old in text:
                text = text.replace(old, old.replace(f"  {slug} version", usage_line + "\n  " + slug + " version", 1))
                break

    path.write_text(text)
    return True


def write_network_go(cmd_dir: Path) -> bool:
    path = cmd_dir / "network.go"
    if path.exists():
        return False
    path.write_text(NETWORK_BODY)
    return True


def main() -> None:
    changed = 0
    for main_path in sorted(ROOT.glob("*-cli/cmd/*/main.go")):
        slug = main_path.parent.name
        c1 = patch_main(main_path, slug)
        c2 = write_network_go(main_path.parent)
        if c1 or c2:
            changed += 1
    print(f"patched {changed} CLIs")


if __name__ == "__main__":
    main()
