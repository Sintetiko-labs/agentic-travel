#!/usr/bin/env python3
"""Scaffold agentic-travel CLIs from brand group definitions."""

from __future__ import annotations

import json
import os
import re
import subprocess
import sys
import textwrap
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent

# Each group: slug, display name, category (hotel|airline), base_url, brands[]
GROUPS: list[dict] = [
    # --- Hotels (Spanish & international) ---
    {"slug": "melia", "name": "Meliá", "cat": "hotel", "url": "https://www.melia.com",
     "brands": ["Meliá Hotels International", "Meliá", "Gran Meliá", "ME by Meliá", "The Meliá Collection", "Paradisus", "INNSiDE by Meliá", "Sol by Meliá", "ZEL"]},
    {"slug": "barcelo", "name": "Barceló", "cat": "hotel", "url": "https://www.barcelo.com",
     "brands": ["Barceló Hotel Group", "Barceló Hotels & Resorts", "Royal Hideaway", "Occidental Hotels & Resorts", "Allegro Hotels"]},
    {"slug": "riu", "name": "RIU", "cat": "hotel", "url": "https://www.riu.com", "brands": ["RIU Hotels & Resorts"]},
    {"slug": "iberostar", "name": "Iberostar", "cat": "hotel", "url": "https://www.iberostar.com",
     "brands": ["Iberostar", "Iberostar Selection", "Iberostar Grand"]},
    {"slug": "nh", "name": "NH Hotel Group", "cat": "hotel", "url": "https://www.nh-hotels.com",
     "brands": ["NH Hotel Group", "NH Hotels", "NH Collection", "nhow"]},
    {"slug": "minor", "name": "Minor Hotels", "cat": "hotel", "url": "https://www.minorhotels.com",
     "brands": ["Avani", "Tivoli", "Minor Hotels"]},
    {"slug": "eurostars", "name": "Eurostars", "cat": "hotel", "url": "https://www.eurostarshotels.com",
     "brands": ["Eurostars Hotel Company", "Eurostars Hotels", "Exe Hotels", "Ikonik Hotels", "Áurea Hotels", "Tandem Suites"]},
    {"slug": "hotusa", "name": "Hotusa", "cat": "hotel", "url": "https://www.hotusa.com",
     "brands": ["Hotusa", "Crisol Hotels"]},
    {"slug": "h10", "name": "H10", "cat": "hotel", "url": "https://www.h10hotels.com",
     "brands": ["H10 Hotels", "H10", "Ocean by H10"]},
    {"slug": "princess", "name": "Princess Hotels", "cat": "hotel", "url": "https://www.princess-hotels.com", "brands": ["Princess Hotels"]},
    {"slug": "catalonia", "name": "Catalonia Hotels", "cat": "hotel", "url": "https://www.cataloniahotels.com", "brands": ["Catalonia Hotels & Resorts"]},
    {"slug": "vincci", "name": "Vincci", "cat": "hotel", "url": "https://www.vinccihoteles.com", "brands": ["Vincci Hoteles"]},
    {"slug": "silken", "name": "Silken", "cat": "hotel", "url": "https://www.hoteles-silken.com", "brands": ["Silken Hoteles"]},
    {"slug": "sercotel", "name": "Sercotel", "cat": "hotel", "url": "https://www.sercotelhoteles.com", "brands": ["Sercotel"]},
    {"slug": "roommate", "name": "Room Mate", "cat": "hotel", "url": "https://www.room-matehotels.com", "brands": ["Room Mate Hotels"]},
    {"slug": "onlyyou", "name": "Only YOU", "cat": "hotel", "url": "https://www.onlyyouhotels.com", "brands": ["Only YOU Hotels"]},
    {"slug": "palladium", "name": "Palladium", "cat": "hotel", "url": "https://www.palladiumhotelgroup.com",
     "brands": ["Palladium Hotel Group", "Ushuaïa Ibiza Beach Hotel", "Hard Rock Hotel Ibiza", "TRS Hotels", "Grand Palladium", "Palladium Hotels"]},
    {"slug": "bless", "name": "BLESS Collection", "cat": "hotel", "url": "https://www.blesscollectionhotels.com", "brands": ["BLESS Collection Hotels"]},
    {"slug": "pinero", "name": "Grupo Piñero", "cat": "hotel", "url": "https://www.bahia-principe.com",
     "brands": ["Fiesta Hotels & Resorts", "Grupo Piñero", "Bahia Principe"]},
    {"slug": "senator", "name": "Senator", "cat": "hotel", "url": "https://www.senator.es",
     "brands": ["Senator Hotels & Resorts", "Playa Senator"]},
    {"slug": "hipotels", "name": "Hipotels", "cat": "hotel", "url": "https://www.hipotels.com", "brands": ["Hipotels"]},
    {"slug": "lopesan", "name": "Lopesan", "cat": "hotel", "url": "https://www.lopesan.com",
     "brands": ["Lopesan Hotel Group", "Abora by Lopesan", "Lopesan Hotels", "Lopesan Collection"]},
    {"slug": "seaside", "name": "Seaside Collection", "cat": "hotel", "url": "https://www.seaside-collection.com", "brands": ["Seaside Collection"]},
    {"slug": "belive", "name": "Be Live", "cat": "hotel", "url": "https://www.belivehotels.com", "brands": ["Be Live Hotels"]},
    {"slug": "globales", "name": "Globales", "cat": "hotel", "url": "https://www.globales.com",
     "brands": ["Globales Hotels", "Hoteles Globales"]},
    {"slug": "grupotel", "name": "Grupotel", "cat": "hotel", "url": "https://www.grupotel.com", "brands": ["Grupotel"]},
    {"slug": "garden", "name": "Garden Hotels", "cat": "hotel", "url": "https://www.gardenhotels.com", "brands": ["Garden Hotels"]},
    {"slug": "zafiro", "name": "Zafiro", "cat": "hotel", "url": "https://www.zafirohotels.com", "brands": ["Zafiro Hotels"]},
    {"slug": "viva", "name": "Viva Hotels", "cat": "hotel", "url": "https://www.vivahotels.com", "brands": ["Viva Hotels"]},
    {"slug": "protur", "name": "Protur", "cat": "hotel", "url": "https://www.protur-hotels.com", "brands": ["Protur Hotels"]},
    {"slug": "fergus", "name": "Fergus", "cat": "hotel", "url": "https://www.fergushotels.com", "brands": ["Fergus Hotels"]},
    {"slug": "tent", "name": "Tent Hotels", "cat": "hotel", "url": "https://www.tenthotels.com", "brands": ["Tent Hotels"]},
    {"slug": "iberik", "name": "Iberik", "cat": "hotel", "url": "https://www.iberikhotels.com", "brands": ["Iberik Hoteles"]},
    {"slug": "servigroup", "name": "Servigroup", "cat": "hotel", "url": "https://www.servigroup.com", "brands": ["Hoteles Servigroup"]},
    {"slug": "medplaya", "name": "MedPlaya", "cat": "hotel", "url": "https://www.medplaya.com", "brands": ["MedPlaya"]},
    {"slug": "besthotels", "name": "Best Hotels", "cat": "hotel", "url": "https://www.besthotels.es", "brands": ["Best Hotels"]},
    {"slug": "alegria", "name": "Alegria", "cat": "hotel", "url": "https://www.alegriahotels.com", "brands": ["Alegria Hotels"]},
    {"slug": "htop", "name": "HTop", "cat": "hotel", "url": "https://www.htophotels.com", "brands": ["HTop Hotels"]},
    {"slug": "guitart", "name": "Guitart", "cat": "hotel", "url": "https://www.guitarthotels.com", "brands": ["Guitart Hotels"]},
    {"slug": "evenia", "name": "Evenia", "cat": "hotel", "url": "https://www.eveniahotels.com", "brands": ["Evenia Hotels"]},
    {"slug": "sbhotels", "name": "SB Hotels", "cat": "hotel", "url": "https://www.sbhotels.com", "brands": ["SB Hotels"]},
    {"slug": "ilunion", "name": "Ilunion", "cat": "hotel", "url": "https://www.ilunionhotels.com", "brands": ["Ilunion Hotels"]},
    {"slug": "paradores", "name": "Paradores", "cat": "hotel", "url": "https://www.parador.es", "brands": ["Paradores"]},
    {"slug": "soho", "name": "Soho Boutique", "cat": "hotel", "url": "https://www.sohohoteles.com", "brands": ["Soho Boutique Hotels"]},
    {"slug": "casual", "name": "Casual Hoteles", "cat": "hotel", "url": "https://www.casualhoteles.com", "brands": ["Casual Hoteles"]},
    {"slug": "petitpalace", "name": "Petit Palace", "cat": "hotel", "url": "https://www.petitpalace.com", "brands": ["Petit Palace"]},
    {"slug": "hightech", "name": "High Tech Hotels", "cat": "hotel", "url": "https://www.hthoteles.com", "brands": ["High Tech Hotels"]},
    {"slug": "oneshot", "name": "One Shot", "cat": "hotel", "url": "https://www.oneshothotels.com", "brands": ["One Shot Hotels"]},
    {"slug": "umusic", "name": "UMusic Hotels", "cat": "hotel", "url": "https://www.umusichotels.com", "brands": ["UMusic Hotels"]},
    {"slug": "abba", "name": "Abba Hoteles", "cat": "hotel", "url": "https://www.abbahoteles.com", "brands": ["Abba Hoteles"]},
    {"slug": "zenit", "name": "Zenit", "cat": "hotel", "url": "https://www.zenithoteles.com", "brands": ["Zenit Hoteles"]},
    {"slug": "vp", "name": "VP Hoteles", "cat": "hotel", "url": "https://www.vp-hoteles.com", "brands": ["VP Hoteles"]},
    {"slug": "derby", "name": "Derby Hotels", "cat": "hotel", "url": "https://www.derbyhotels.com", "brands": ["Derby Hotels Collection"]},
    {"slug": "alma", "name": "Alma Hotels", "cat": "hotel", "url": "https://www.almahotels.com", "brands": ["Alma Hotels"]},
    {"slug": "hospes", "name": "Hospes", "cat": "hotel", "url": "https://www.hospes.com", "brands": ["Hospes Hotels"]},
    {"slug": "unico", "name": "Único Hotels", "cat": "hotel", "url": "https://www.unicohotels.com", "brands": ["Único Hotels"]},
    {"slug": "coolrooms", "name": "CoolRooms", "cat": "hotel", "url": "https://www.coolrooms.com", "brands": ["CoolRooms Hotels"]},
    {"slug": "castillatermal", "name": "Castilla Termal", "cat": "hotel", "url": "https://www.castillatermal.com", "brands": ["Castilla Termal"]},
    {"slug": "eurobuilding", "name": "Eurobuilding", "cat": "hotel", "url": "https://www.eurobuilding.es", "brands": ["Eurobuilding"]},
    {"slug": "center", "name": "Hoteles Center", "cat": "hotel", "url": "https://www.hotelescenter.com", "brands": ["Hoteles Center"]},
    {"slug": "santos", "name": "Hoteles Santos", "cat": "hotel", "url": "https://www.hoteles-santos.com", "brands": ["Hoteles Santos"]},
    {"slug": "elba", "name": "Hoteles Elba", "cat": "hotel", "url": "https://www.hoteleselba.com", "brands": ["Hoteles Elba"]},
    {"slug": "poseidon", "name": "Hoteles Poseidón", "cat": "hotel", "url": "https://www.hoteles-poseidon.com", "brands": ["Hoteles Poseidón"]},
    {"slug": "rh", "name": "Hoteles RH", "cat": "hotel", "url": "https://www.rhhotels.com", "brands": ["Hoteles RH"]},
    {"slug": "magic", "name": "Magic Costa Blanca", "cat": "hotel", "url": "https://www.magiccostablanca.com", "brands": ["Magic Costa Blanca"]},
    {"slug": "porthotels", "name": "Port Hotels", "cat": "hotel", "url": "https://www.porthotels.es", "brands": ["Port Hotels"]},
    {"slug": "ona", "name": "Ona Hotels", "cat": "hotel", "url": "https://www.onahotels.com",
     "brands": ["Ona Hotels", "Ona Hotels & Apartments"]},
    {"slug": "pierrevacances", "name": "Pierre & Vacances", "cat": "hotel", "url": "https://www.pierreetvacances.com",
     "brands": ["Pierre & Vacances", "Pierre & Vacances España", "Apartamentos Pierre & Vacances"]},
    {"slug": "smartrental", "name": "SmartRental", "cat": "hotel", "url": "https://www.smartrental.com", "brands": ["SmartRental"]},
    {"slug": "libere", "name": "Líbere", "cat": "hotel", "url": "https://www.liberehospitality.com", "brands": ["Líbere Hospitality"]},
    {"slug": "numa", "name": "Numa", "cat": "hotel", "url": "https://www.numastays.com", "brands": ["Numa"]},
    {"slug": "sonder", "name": "Sonder", "cat": "hotel", "url": "https://www.sonder.com", "brands": ["Sonder"]},
    {"slug": "limehome", "name": "Limehome", "cat": "hotel", "url": "https://www.limehome.com", "brands": ["Limehome"]},
    {"slug": "bbhotels", "name": "B&B Hotels", "cat": "hotel", "url": "https://www.hotel-bb.com", "brands": ["B&B Hotels"]},
    {"slug": "travelodge", "name": "Travelodge", "cat": "hotel", "url": "https://www.travelodge.co.uk", "brands": ["Travelodge"]},
    {"slug": "easyhotel", "name": "easyHotel", "cat": "hotel", "url": "https://www.easyhotel.com", "brands": ["easyHotel"]},
    {"slug": "accor", "name": "Accor", "cat": "hotel", "url": "https://all.accor.com",
     "brands": ["Ibis", "Ibis Budget", "Ibis Styles", "Novotel", "Mercure", "Pullman", "Sofitel", "MGallery", "Fairmont", "Raffles", "Accor"]},
    {"slug": "marriott", "name": "Marriott", "cat": "hotel", "url": "https://www.marriott.com",
     "brands": ["Marriott", "Marriott Hotels", "JW Marriott", "The Ritz-Carlton", "St. Regis", "W Hotels", "Edition", "Luxury Collection", "Westin", "Sheraton", "Le Méridien", "Renaissance Hotels", "Autograph Collection", "Tribute Portfolio", "AC Hotels", "AC Hotels by Marriott", "Aloft", "Moxy", "Courtyard by Marriott", "Residence Inn"]},
    {"slug": "hilton", "name": "Hilton", "cat": "hotel", "url": "https://www.hilton.com",
     "brands": ["Hilton", "Hilton Hotels & Resorts", "Conrad", "Waldorf Astoria", "DoubleTree by Hilton", "Canopy by Hilton", "Curio Collection", "Hampton by Hilton"]},
    {"slug": "hyatt", "name": "Hyatt", "cat": "hotel", "url": "https://www.hyatt.com",
     "brands": ["Hyatt", "Grand Hyatt", "Hyatt Regency", "Hyatt Centric", "Thompson Hotels", "Andaz", "Alua Hotels", "Dreams Resorts", "Secrets Resorts", "Zoëtry", "Inclusive Collection"]},
    {"slug": "ihg", "name": "IHG", "cat": "hotel", "url": "https://www.ihg.com",
     "brands": ["IHG Hotels & Resorts", "InterContinental", "Kimpton", "Crowne Plaza", "Holiday Inn", "Holiday Inn Express", "Hotel Indigo", "Six Senses", "Vignette Collection"]},
    {"slug": "radisson", "name": "Radisson", "cat": "hotel", "url": "https://www.radissonhotels.com",
     "brands": ["Radisson Hotel Group", "Radisson Blu", "Radisson RED", "Radisson Collection", "Park Inn by Radisson"]},
    {"slug": "leonardo", "name": "Leonardo Hotels", "cat": "hotel", "url": "https://www.leonardo-hotels.com",
     "brands": ["Leonardo Hotels", "NYX Hotels", "Leonardo Royal", "Leonardo Boutique"]},
    {"slug": "wyndham", "name": "Wyndham", "cat": "hotel", "url": "https://www.wyndhamhotels.com",
     "brands": ["Wyndham Hotels & Resorts", "Ramada", "Wyndham", "Tryp", "Dolce by Wyndham"]},
    {"slug": "bestwestern", "name": "Best Western", "cat": "hotel", "url": "https://www.bestwestern.com",
     "brands": ["Best Western", "Best Western Plus", "Best Western Premier", "BWH Hotel Group"]},
    {"slug": "preferred", "name": "Preferred Hotels", "cat": "hotel", "url": "https://www.preferredhotels.com", "brands": ["Preferred Hotels & Resorts"]},
    {"slug": "lhw", "name": "Leading Hotels", "cat": "hotel", "url": "https://www.lhw.com", "brands": ["Leading Hotels of the World"]},
    {"slug": "slh", "name": "Small Luxury Hotels", "cat": "hotel", "url": "https://www.slh.com", "brands": ["Small Luxury Hotels of the World"]},
    {"slug": "relaischateaux", "name": "Relais & Châteaux", "cat": "hotel", "url": "https://www.relaischateaux.com", "brands": ["Relais & Châteaux"]},
    {"slug": "designhotels", "name": "Design Hotels", "cat": "hotel", "url": "https://www.designhotels.com", "brands": ["Design Hotels"]},
    {"slug": "mandarin", "name": "Mandarin Oriental", "cat": "hotel", "url": "https://www.mandarinoriental.com", "brands": ["Mandarin Oriental"]},
    {"slug": "fourseasons", "name": "Four Seasons", "cat": "hotel", "url": "https://www.fourseasons.com", "brands": ["Four Seasons"]},
    {"slug": "rosewood", "name": "Rosewood", "cat": "hotel", "url": "https://www.rosewoodhotels.com", "brands": ["Rosewood"]},
    {"slug": "belmond", "name": "Belmond", "cat": "hotel", "url": "https://www.belmond.com", "brands": ["Belmond"]},
    {"slug": "aman", "name": "Aman", "cat": "hotel", "url": "https://www.aman.com", "brands": ["Aman"]},
    {"slug": "nobu", "name": "Nobu Hotels", "cat": "hotel", "url": "https://www.nobuhotels.com", "brands": ["Nobu Hotels"]},
    {"slug": "virgin", "name": "Virgin Hotels", "cat": "hotel", "url": "https://www.virginhotels.com", "brands": ["Virgin Hotels"]},
    {"slug": "citizenm", "name": "citizenM", "cat": "hotel", "url": "https://www.citizenm.com", "brands": ["citizenM"]},
    {"slug": "mamashelter", "name": "Mama Shelter", "cat": "hotel", "url": "https://www.mamashelter.com", "brands": ["Mama Shelter"]},
    {"slug": "hoxton", "name": "The Hoxton", "cat": "hotel", "url": "https://www.thehoxton.com", "brands": ["The Hoxton"]},
    {"slug": "25hours", "name": "25hours Hotels", "cat": "hotel", "url": "https://www.25hours-hotels.com", "brands": ["25hours Hotels"]},
    {"slug": "ruby", "name": "Ruby Hotels", "cat": "hotel", "url": "https://www.ruby-hotels.com", "brands": ["Ruby Hotels"]},
    {"slug": "zoku", "name": "Zoku", "cat": "hotel", "url": "https://www.livezoku.com", "brands": ["Zoku"]},
    {"slug": "locke", "name": "Locke", "cat": "hotel", "url": "https://www.lockeliving.com", "brands": ["Locke"]},
    {"slug": "bypillow", "name": "ByPillow", "cat": "hotel", "url": "https://www.bypillow.com", "brands": ["ByPillow"]},
    {"slug": "axel", "name": "Axel Hotels", "cat": "hotel", "url": "https://www.axelhotels.com", "brands": ["Axel Hotels"]},
    {"slug": "generator", "name": "Generator Hostels", "cat": "hotel", "url": "https://www.staygenerator.com", "brands": ["Generator Hostels"]},
    {"slug": "toc", "name": "TOC Hostels", "cat": "hotel", "url": "https://www.tochostels.com", "brands": ["TOC Hostels"]},
    {"slug": "latroupe", "name": "Latroupe", "cat": "hotel", "url": "https://www.latroupe.com", "brands": ["Latroupe"]},
    {"slug": "safestay", "name": "Safestay", "cat": "hotel", "url": "https://www.safestay.com", "brands": ["Safestay"]},
    {"slug": "stchristophers", "name": "St Christopher's", "cat": "hotel", "url": "https://www.st-christophers.co.uk", "brands": ["St Christopher's Inns Iberia"]},
    # --- Airlines ---
    {"slug": "iberiaexpress", "name": "Iberia Express", "cat": "airline", "url": "https://www.iberiaexpress.com", "brands": ["Iberia Express"]},
    {"slug": "vueling", "name": "Vueling", "cat": "airline", "url": "https://www.vueling.com", "brands": ["Vueling"]},
    {"slug": "aireuropa", "name": "Air Europa", "cat": "airline", "url": "https://www.aireuropa.com", "brands": ["Air Europa"]},
    {"slug": "ryanair", "name": "Ryanair", "cat": "airline", "url": "https://www.ryanair.com", "brands": ["Ryanair"]},
    {"slug": "easyjet", "name": "easyJet", "cat": "airline", "url": "https://www.easyjet.com", "brands": ["easyJet"]},
    {"slug": "wizzair", "name": "Wizz Air", "cat": "airline", "url": "https://wizzair.com", "brands": ["Wizz Air"]},
    {"slug": "volotea", "name": "Volotea", "cat": "airline", "url": "https://www.volotea.com", "brands": ["Volotea"]},
    {"slug": "binter", "name": "Binter", "cat": "airline", "url": "https://www.bintercanarias.com", "brands": ["Binter"]},
    {"slug": "canaryfly", "name": "Canaryfly", "cat": "airline", "url": "https://www.canaryfly.es", "brands": ["Canaryfly"]},
    {"slug": "level", "name": "Level", "cat": "airline", "url": "https://www.flylevel.com", "brands": ["Level"]},
    {"slug": "plusultra", "name": "Plus Ultra", "cat": "airline", "url": "https://www.plusultra.com", "brands": ["Plus Ultra Líneas Aéreas"]},
    {"slug": "world2fly", "name": "World2Fly", "cat": "airline", "url": "https://www.world2fly.com", "brands": ["World2Fly"]},
    {"slug": "iberojet", "name": "Iberojet", "cat": "airline", "url": "https://www.iberojet.es", "brands": ["Iberojet"]},
    {"slug": "privilegestyle", "name": "Privilege Style", "cat": "airline", "url": "https://www.privilegestyle.com", "brands": ["Privilege Style"]},
    {"slug": "airnostrum", "name": "Air Nostrum", "cat": "airline", "url": "https://www.airnostrum.es", "brands": ["Air Nostrum"]},
    {"slug": "swiftair", "name": "Swiftair", "cat": "airline", "url": "https://www.swiftair.com", "brands": ["Swiftair"]},
    {"slug": "albastar", "name": "Albastar", "cat": "airline", "url": "https://www.albastar.es", "brands": ["Albastar"]},
    {"slug": "tui", "name": "TUI", "cat": "airline", "url": "https://www.tui.co.uk", "brands": ["TUI Airways", "TUI fly"]},
    {"slug": "jet2", "name": "Jet2", "cat": "airline", "url": "https://www.jet2.com", "brands": ["Jet2"]},
    {"slug": "norwegian", "name": "Norwegian", "cat": "airline", "url": "https://www.norwegian.com", "brands": ["Norwegian"]},
    {"slug": "lufthansagroup", "name": "Lufthansa Group", "cat": "airline", "url": "https://www.lufthansa.com",
     "brands": ["Lufthansa", "Lufthansa City Airlines", "Discover Airlines", "Swiss", "Austrian Airlines", "Brussels Airlines", "Eurowings"]},
    {"slug": "airfranceklm", "name": "Air France-KLM", "cat": "airline", "url": "https://www.airfrance.com",
     "brands": ["Air France", "KLM", "Transavia"]},
    {"slug": "britishairways", "name": "British Airways", "cat": "airline", "url": "https://www.britishairways.com",
     "brands": ["British Airways", "BA CityFlyer"]},
    {"slug": "aerlingus", "name": "Aer Lingus", "cat": "airline", "url": "https://www.aerlingus.com", "brands": ["Aer Lingus"]},
    {"slug": "tap", "name": "TAP Air Portugal", "cat": "airline", "url": "https://www.flytap.com", "brands": ["TAP Air Portugal"]},
    {"slug": "ita", "name": "ITA Airways", "cat": "airline", "url": "https://www.ita-airways.com", "brands": ["ITA Airways"]},
    {"slug": "sas", "name": "SAS", "cat": "airline", "url": "https://www.flysas.com", "brands": ["SAS"]},
    {"slug": "finnair", "name": "Finnair", "cat": "airline", "url": "https://www.finnair.com", "brands": ["Finnair"]},
    {"slug": "lot", "name": "LOT Polish Airlines", "cat": "airline", "url": "https://www.lot.com", "brands": ["LOT Polish Airlines"]},
    {"slug": "czechairlines", "name": "Czech Airlines", "cat": "airline", "url": "https://www.csa.cz", "brands": ["Czech Airlines"]},
    {"slug": "smartwings", "name": "Smartwings", "cat": "airline", "url": "https://www.smartwings.com", "brands": ["Smartwings"]},
    {"slug": "aegean", "name": "Aegean", "cat": "airline", "url": "https://www.aegeanair.com",
     "brands": ["Aegean Airlines", "Olympic Air"]},
    {"slug": "croatiaairlines", "name": "Croatia Airlines", "cat": "airline", "url": "https://www.croatiaairlines.com", "brands": ["Croatia Airlines"]},
    {"slug": "airserbia", "name": "Air Serbia", "cat": "airline", "url": "https://www.airserbia.com", "brands": ["Air Serbia"]},
    {"slug": "turkish", "name": "Turkish Airlines Group", "cat": "airline", "url": "https://www.turkishairlines.com",
     "brands": ["Turkish Airlines", "Pegasus Airlines", "SunExpress", "AnadoluJet / AJet"]},
    {"slug": "emirates", "name": "Emirates", "cat": "airline", "url": "https://www.emirates.com", "brands": ["Emirates"]},
    {"slug": "qatar", "name": "Qatar Airways", "cat": "airline", "url": "https://www.qatarairways.com", "brands": ["Qatar Airways"]},
    {"slug": "etihad", "name": "Etihad Airways", "cat": "airline", "url": "https://www.etihad.com", "brands": ["Etihad Airways"]},
    {"slug": "saudia", "name": "Saudia", "cat": "airline", "url": "https://www.saudia.com", "brands": ["Saudia"]},
    {"slug": "royaljordanian", "name": "Royal Jordanian", "cat": "airline", "url": "https://www.rj.com", "brands": ["Royal Jordanian"]},
    {"slug": "kuwaitairways", "name": "Kuwait Airways", "cat": "airline", "url": "https://www.kuwaitairways.com", "brands": ["Kuwait Airways"]},
    {"slug": "gulfair", "name": "Gulf Air", "cat": "airline", "url": "https://www.gulfair.com", "brands": ["Gulf Air"]},
    {"slug": "egyptair", "name": "Egyptair", "cat": "airline", "url": "https://www.egyptair.com", "brands": ["Egyptair"]},
    {"slug": "royalairmaroc", "name": "Royal Air Maroc", "cat": "airline", "url": "https://www.royalairmaroc.com", "brands": ["Royal Air Maroc"]},
    {"slug": "airarabia", "name": "Air Arabia", "cat": "airline", "url": "https://www.airarabia.com",
     "brands": ["Air Arabia", "Air Arabia Maroc"]},
    {"slug": "tunisair", "name": "Tunisair", "cat": "airline", "url": "https://www.tunisair.com", "brands": ["Tunisair"]},
    {"slug": "nouvelair", "name": "Nouvelair", "cat": "airline", "url": "https://www.nouvelair.com", "brands": ["Nouvelair"]},
    {"slug": "airalgerie", "name": "Air Algérie", "cat": "airline", "url": "https://www.airalgerie.dz", "brands": ["Air Algerie"]},
    {"slug": "caboverde", "name": "Cabo Verde Airlines", "cat": "airline", "url": "https://www.caboverdeairlines.com", "brands": ["Cabo Verde Airlines"]},
    {"slug": "ethiopian", "name": "Ethiopian Airlines", "cat": "airline", "url": "https://www.ethiopianairlines.com", "brands": ["Ethiopian Airlines"]},
    {"slug": "kenyaairways", "name": "Kenya Airways", "cat": "airline", "url": "https://www.kenya-airways.com", "brands": ["Kenya Airways"]},
    {"slug": "united", "name": "United Airlines", "cat": "airline", "url": "https://www.united.com", "brands": ["United Airlines"]},
    {"slug": "american", "name": "American Airlines", "cat": "airline", "url": "https://www.aa.com", "brands": ["American Airlines"]},
    {"slug": "delta", "name": "Delta Air Lines", "cat": "airline", "url": "https://www.delta.com", "brands": ["Delta Air Lines"]},
    {"slug": "aircanada", "name": "Air Canada", "cat": "airline", "url": "https://www.aircanada.com", "brands": ["Air Canada"]},
    {"slug": "airtransat", "name": "Air Transat", "cat": "airline", "url": "https://www.airtransat.com", "brands": ["Air Transat"]},
    {"slug": "westjet", "name": "WestJet", "cat": "airline", "url": "https://www.westjet.com", "brands": ["WestJet"]},
    {"slug": "aeromexico", "name": "Aeroméxico", "cat": "airline", "url": "https://www.aeromexico.com", "brands": ["Aeroméxico"]},
    {"slug": "avianca", "name": "Avianca", "cat": "airline", "url": "https://www.avianca.com", "brands": ["Avianca"]},
    {"slug": "latam", "name": "LATAM", "cat": "airline", "url": "https://www.latam.com",
     "brands": ["LATAM Airlines", "LATAM Brasil", "LATAM Chile"]},
    {"slug": "aerolineas", "name": "Aerolíneas Argentinas", "cat": "airline", "url": "https://www.aerolineas.com.ar", "brands": ["Aerolíneas Argentinas"]},
    {"slug": "copa", "name": "Copa Airlines", "cat": "airline", "url": "https://www.copaair.com", "brands": ["Copa Airlines"]},
    {"slug": "airchina", "name": "Air China", "cat": "airline", "url": "https://www.airchina.com", "brands": ["Air China"]},
    {"slug": "chinaeastern", "name": "China Eastern", "cat": "airline", "url": "https://www.ceair.com", "brands": ["China Eastern"]},
    {"slug": "chinasouthern", "name": "China Southern", "cat": "airline", "url": "https://www.csair.com", "brands": ["China Southern"]},
    {"slug": "hainan", "name": "Hainan Airlines", "cat": "airline", "url": "https://www.hainanairlines.com", "brands": ["Hainan Airlines"]},
    {"slug": "cathaypacific", "name": "Cathay Pacific", "cat": "airline", "url": "https://www.cathaypacific.com", "brands": ["Cathay Pacific"]},
    {"slug": "koreanair", "name": "Korean Air", "cat": "airline", "url": "https://www.koreanair.com", "brands": ["Korean Air"]},
    {"slug": "asiana", "name": "Asiana Airlines", "cat": "airline", "url": "https://www.flyasiana.com", "brands": ["Asiana Airlines"]},
    {"slug": "singaporeairlines", "name": "Singapore Airlines", "cat": "airline", "url": "https://www.singaporeair.com", "brands": ["Singapore Airlines"]},
    {"slug": "thaiairways", "name": "Thai Airways", "cat": "airline", "url": "https://www.thaiairways.com", "brands": ["Thai Airways"]},
    {"slug": "vietnamairlines", "name": "Vietnam Airlines", "cat": "airline", "url": "https://www.vietnamairlines.com", "brands": ["Vietnam Airlines"]},
    {"slug": "qantas", "name": "Qantas", "cat": "airline", "url": "https://www.qantas.com", "brands": ["Qantas"]},
    {"slug": "elal", "name": "El Al", "cat": "airline", "url": "https://www.elal.com", "brands": ["El Al"]},
    {"slug": "icelandair", "name": "Icelandair", "cat": "airline", "url": "https://www.icelandair.com", "brands": ["Icelandair"]},
    {"slug": "play", "name": "PLAY Airlines", "cat": "airline", "url": "https://www.flyplay.com", "brands": ["PLAY Airlines"]},
    {"slug": "norse", "name": "Norse Atlantic Airways", "cat": "airline", "url": "https://www.flynorse.com", "brands": ["Norse Atlantic Airways"]},
    {"slug": "condor", "name": "Condor", "cat": "airline", "url": "https://www.condor.com", "brands": ["Condor"]},
    {"slug": "corendon", "name": "Corendon Airlines", "cat": "airline", "url": "https://www.corendonairlines.com", "brands": ["Corendon Airlines"]},
    {"slug": "freebird", "name": "Freebird Airlines", "cat": "airline", "url": "https://www.freebirdairlines.com", "brands": ["Freebird Airlines"]},
    {"slug": "enterair", "name": "Enter Air", "cat": "airline", "url": "https://www.enterair.pl", "brands": ["Enter Air"]},
    {"slug": "neos", "name": "Neos", "cat": "airline", "url": "https://www.neosair.it", "brands": ["Neos"]},
    {"slug": "wamos", "name": "Wamos Air", "cat": "airline", "url": "https://www.wamosair.com", "brands": ["Wamos Air"]},
]


def env_prefix(slug: str) -> str:
    return slug.upper().replace("-", "_")


def write(path: Path, content: str) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(content, encoding="utf-8")


def gen_hotel_cli(g: dict) -> None:
    slug = g["slug"]
    name = g["name"]
    mod = f"github.com/fbelchi/{slug}-cli"
    pkg = slug.replace("-", "")
    if pkg[0].isdigit():
        pkg = "x" + pkg
    brands_go = ",\n\t\t".join(f'"{b}"' for b in g["brands"])
    brand_flag = len(g["brands"]) > 1

    write(ROOT / f"{slug}-cli/go.mod", f"""module {mod}

go 1.26

require github.com/fbelchi/travelkit v0.0.0

replace github.com/fbelchi/travelkit => ../travelkit
""")

    write(ROOT / f"{slug}-cli/.gitignore", f"/{slug}\n*.exe\n")

    write(ROOT / f"{slug}-cli/internal/client/client.go", f"""package client

import (
\ttkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "{g['url']}"

// Client talks to {name} public endpoints.
type Client struct {{
\t*tkbase.Client
\tBrand string
}}

// Brands supported by this CLI (shared parent API).
var Brands = []string{{
\t\t{brands_go},
}}

func New(brand string) *Client {{
\treturn &Client{{Client: tkbase.New(BaseURL, "{slug}"), Brand: brand}}
}}
""")

    write(ROOT / f"{slug}-cli/internal/client/types.go", """package client

import tktypes "github.com/fbelchi/travelkit/types"

type HotelSearchResult = tktypes.HotelSearchResult
type HotelHit = tktypes.HotelHit
type HotelView = tktypes.HotelView
type AvailSummary = tktypes.AvailSummary
type PriceInfo = tktypes.PriceInfo
""")

    write(ROOT / f"{slug}-cli/internal/client/errors.go", f"""package client

import "fmt"

type APIError struct {{
\tStatus int
\tBody   string
}}

func (e *APIError) Error() string {{
\treturn fmt.Sprintf("{slug} api: HTTP %d: %s", e.Status, e.Body)
}}
""")

    write(ROOT / f"{slug}-cli/internal/client/search.go", f"""package client

import "fmt"

// Search runs hotel search (TODO: implement for {name}).
func (c *Client) Search(query string, page, pageSize int) (*HotelSearchResult, error) {{
\tif page < 1 {{
\t\tpage = 1
\t}}
\tif pageSize < 1 {{
\t\tpageSize = 24
\t}}
\t_ = c
\treturn nil, fmt.Errorf("search not yet implemented for {name} — see README and internal/client/search.go TODO")
}}
""")

    write(ROOT / f"{slug}-cli/internal/client/read.go", f"""package client

import "fmt"

// Read returns hotel detail (stub).
func (c *Client) Read(idOrURL string) (*HotelView, error) {{
\treturn nil, fmt.Errorf("read not yet implemented for {name} (id=%q)", idOrURL)
}}
""")

    write(ROOT / f"{slug}-cli/internal/client/availability.go", f"""package client

import "fmt"

// Availability checks room availability (stub).
func (c *Client) Availability(hotelID, checkIn, checkOut string, guests, rooms int) (*AvailSummary, error) {{
\treturn nil, fmt.Errorf("availability not yet implemented for {name} (hotel=%q)", hotelID)
}}
""")

    brand_help = ""
    brand_flag_code = ""
    if brand_flag:
        brand_help = f"\n  {slug} search [--json] [--brand BRAND] [--limit N] <destination...>"
        brand_flag_code = """
\tbrand := fs.String("brand", "", "sub-brand (see --help brands)")
"""
    else:
        brand_help = ""

    write(ROOT / f"{slug}-cli/cmd/{slug}/main.go", f"""// Command {slug} is an unofficial, agent-friendly CLI for {name}.
package main

import (
\t"fmt"
\t"os"
)

var version = "dev"

func main() {{
\tif len(os.Args) < 2 {{
\t\tusage()
\t\tos.Exit(2)
\t}}
\tvar err error
\tswitch os.Args[1] {{
\tcase "search":
\t\terr = cmdSearch(os.Args[2:])
\tcase "read":
\t\terr = cmdRead(os.Args[2:])
\tcase "availability":
\t\terr = cmdAvailability(os.Args[2:])
\tcase "brands":
\t\tcmdBrands()
\tcase "version", "--version", "-v":
\t\tfmt.Println(version)
\tcase "help", "-h", "--help":
\t\tusage()
\tdefault:
\t\tfmt.Fprintf(os.Stderr, "unknown command %q\\n\\n", os.Args[1])
\t\tusage()
\t\tos.Exit(2)
\t}}
\tif err != nil {{
\t\tfmt.Fprintln(os.Stderr, "error:", err)
\t\tos.Exit(1)
\t}}
}}

func usage() {{
\tfmt.Fprintf(os.Stderr, `{slug} — unofficial CLI for {name}

USAGE:
  {slug} search [--json] [--limit N]{' [--brand BRAND]' if brand_flag else ''} <destination...>
  {slug} read [--json]{' [--brand BRAND]' if brand_flag else ''} <id|url>
  {slug} availability [--json]{' [--brand BRAND]' if brand_flag else ''} --check-in DATE --check-out DATE [--guests N] [--rooms N] <hotel-id>
  {slug} brands
  {slug} version | help
`)
}}
""")

    write(ROOT / f"{slug}-cli/cmd/{slug}/cli.go", """package main

import (
	"encoding/json"
	"flag"
	"os"
	"strings"
)

type common struct {
	jsonOut bool
}

func addCommon(fs *flag.FlagSet) *common {
	c := &common{}
	fs.BoolVar(&c.jsonOut, "json", false, "emit JSON to stdout")
	return c
}

func reorderArgs(fs *flag.FlagSet, args []string) []string {
	var flags, positional []string
	for i := 0; i < len(args); i++ {
		a := args[i]
		if a == "--" {
			positional = append(positional, args[i+1:]...)
			break
		}
		if len(a) > 1 && a[0] == '-' {
			flags = append(flags, a)
			name := strings.TrimLeft(a, "-")
			if strings.IndexByte(name, '=') >= 0 {
				continue
			}
			if f := fs.Lookup(name); f != nil {
				if _, ok := f.Value.(interface{ IsBoolFlag() bool }); !ok && i+1 < len(args) {
					flags = append(flags, args[i+1])
					i++
				}
			}
			continue
		}
		positional = append(positional, a)
	}
	out := append(flags, "--")
	return append(out, positional...)
}

func emitJSON(v any) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)
	return enc.Encode(v)
}
""")

    brand_param_search = 'brand := fs.String("brand", "", "sub-brand")\n\t' if brand_flag else ""
    brand_param_read = brand_param_search
    brand_param_avail = brand_param_search
    brand_arg = "*brand" if brand_flag else '""'

    write(ROOT / f"{slug}-cli/cmd/{slug}/search.go", f"""package main

import (
\t"flag"
\t"fmt"
\t"strings"

\t"github.com/fbelchi/{slug}-cli/internal/client"
)

func cmdSearch(args []string) error {{
\tfs := flag.NewFlagSet("search", flag.ExitOnError)
\tcf := addCommon(fs)
\t{brand_param_search}limit := fs.Int("limit", 24, "max results")
\tpage := fs.Int("page", 1, "page number")
\t_ = fs.Parse(reorderArgs(fs, args))
\tif fs.NArg() == 0 {{
\t\treturn fmt.Errorf("usage: {slug} search [flags] <destination...>")
\t}}
\tcl := client.New({brand_arg})
\tquery := strings.Join(fs.Args(), " ")
\tres, err := cl.Search(query, *page, *limit)
\tif err != nil {{
\t\treturn err
\t}}
\tif cf.jsonOut {{
\t\treturn emitJSON(res)
\t}}
\tfmt.Printf("query=%q total=%d page=%d source=%s\\n", res.Query, res.Total, res.Page, res.Source)
\tfor _, h := range res.Hotels {{
\t\tfmt.Printf("  [%s] %s — %s\\n", h.ID, h.Name, h.Price)
\t}}
\treturn nil
}}
""")

    write(ROOT / f"{slug}-cli/cmd/{slug}/read.go", f"""package main

import (
\t"flag"
\t"fmt"

\t"github.com/fbelchi/{slug}-cli/internal/client"
)

func cmdRead(args []string) error {{
\tfs := flag.NewFlagSet("read", flag.ExitOnError)
\tcf := addCommon(fs)
\t{brand_param_read}_ = fs.Parse(reorderArgs(fs, args))
\tif fs.NArg() != 1 {{
\t\treturn fmt.Errorf("usage: {slug} read [flags] <id|url>")
\t}}
\tcl := client.New({brand_arg})
\tpv, err := cl.Read(fs.Arg(0))
\tif err != nil {{
\t\treturn err
\t}}
\tif cf.jsonOut {{
\t\treturn emitJSON(pv)
\t}}
\tfmt.Printf("[%s] %s\\n", pv.ID, pv.Name)
\tif pv.Price.Price != "" {{
\t\tfmt.Printf("  price: %s %s\\n", pv.Price.Price, pv.Price.Currency)
\t}}
\treturn nil
}}
""")

    write(ROOT / f"{slug}-cli/cmd/{slug}/availability.go", f"""package main

import (
\t"flag"
\t"fmt"

\t"github.com/fbelchi/{slug}-cli/internal/client"
)

func cmdAvailability(args []string) error {{
\tfs := flag.NewFlagSet("availability", flag.ExitOnError)
\tcf := addCommon(fs)
\t{brand_param_avail}checkIn := fs.String("check-in", "", "check-in date (YYYY-MM-DD)")
\tcheckOut := fs.String("check-out", "", "check-out date (YYYY-MM-DD)")
\tguests := fs.Int("guests", 2, "number of guests")
\trooms := fs.Int("rooms", 1, "number of rooms")
\t_ = fs.Parse(reorderArgs(fs, args))
\tif fs.NArg() != 1 {{
\t\treturn fmt.Errorf("usage: {slug} availability [flags] <hotel-id>")
\t}}
\tif *checkIn == "" || *checkOut == "" {{
\t\treturn fmt.Errorf("--check-in and --check-out are required")
\t}}
\tcl := client.New({brand_arg})
\tav, err := cl.Availability(fs.Arg(0), *checkIn, *checkOut, *guests, *rooms)
\tif err != nil {{
\t\treturn err
\t}}
\tif cf.jsonOut {{
\t\treturn emitJSON(av)
\t}}
\tfmt.Printf("status=%s from=%s %s\\n", av.Status, av.From, av.Currency)
\treturn nil
}}
""")

    write(ROOT / f"{slug}-cli/cmd/{slug}/brands.go", f"""package main

import (
\t"fmt"

\t"github.com/fbelchi/{slug}-cli/internal/client"
)

func cmdBrands() {{
\tfor _, b := range client.Brands {{
\t\tfmt.Println(b)
\t}}
}}
""")

    brands_md = "\n".join(f"- {b}" for b in g["brands"])
    brand_section = ""
    if brand_flag:
        brand_section = f"""
## Sub-brands

This CLI covers multiple brands sharing the {name} booking API:

{brands_md}

Use `--brand` to select a sub-brand when searching.
"""

    write(ROOT / f"{slug}-cli/README.md", f"""# {name} CLI

Unofficial, agent-friendly CLI for [{name}]({g['url']}).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o {slug} ./cmd/{slug}
```

## Commands

```bash
{slug} search [--json] [--limit N]{' [--brand BRAND]' if brand_flag else ''} <destination>
{slug} read [--json]{' [--brand BRAND]' if brand_flag else ''} <id|url>
{slug} availability [--json]{' [--brand BRAND]' if brand_flag else ''} --check-in DATE --check-out DATE <hotel-id>
{slug} brands
```

## Environment

- `{env_prefix(slug)}_COOKIE` — optional browser cookie when blocked
- `{env_prefix(slug)}_REQUEST_DELAY` — rate limit (e.g. `2s`)
{brand_section}
## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
""")


def gen_airline_cli(g: dict) -> None:
    slug = g["slug"]
    name = g["name"]
    mod = f"github.com/fbelchi/{slug}-cli"
    brands_go = ",\n\t\t".join(f'"{b}"' for b in g["brands"])
    brand_flag = len(g["brands"]) > 1

    write(ROOT / f"{slug}-cli/go.mod", f"""module {mod}

go 1.26

require github.com/fbelchi/travelkit v0.0.0

replace github.com/fbelchi/travelkit => ../travelkit
""")

    write(ROOT / f"{slug}-cli/.gitignore", f"/{slug}\n*.exe\n")

    write(ROOT / f"{slug}-cli/internal/client/client.go", f"""package client

import (
\ttkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "{g['url']}"

type Client struct {{
\t*tkbase.Client
\tBrand string
}}

var Brands = []string{{
\t\t{brands_go},
}}

func New(brand string) *Client {{
\treturn &Client{{Client: tkbase.New(BaseURL, "{slug}"), Brand: brand}}
}}
""")

    write(ROOT / f"{slug}-cli/internal/client/types.go", """package client

import tktypes "github.com/fbelchi/travelkit/types"

type FlightSearchResult = tktypes.FlightSearchResult
type FlightHit = tktypes.FlightHit
type FlightView = tktypes.FlightView
type PriceInfo = tktypes.PriceInfo
""")

    write(ROOT / f"{slug}-cli/internal/client/errors.go", f"""package client

import "fmt"

type APIError struct {{
\tStatus int
\tBody   string
}}

func (e *APIError) Error() string {{
\treturn fmt.Sprintf("{slug} api: HTTP %d: %s", e.Status, e.Body)
}}
""")

    write(ROOT / f"{slug}-cli/internal/client/search.go", f"""package client

import "fmt"

// Search runs flight search (TODO: implement for {name}).
func (c *Client) Search(origin, dest, depart, ret string, page, pageSize int) (*FlightSearchResult, error) {{
\tif page < 1 {{
\t\tpage = 1
\t}}
\tif pageSize < 1 {{
\t\tpageSize = 24
\t}}
\t_ = c
\treturn nil, fmt.Errorf("search not yet implemented for {name} — see README and internal/client/search.go TODO")
}}
""")

    write(ROOT / f"{slug}-cli/internal/client/read.go", f"""package client

import "fmt"

// Read returns flight or fare detail (stub).
func (c *Client) Read(idOrURL string) (*FlightView, error) {{
\treturn nil, fmt.Errorf("read not yet implemented for {name} (id=%q)", idOrURL)
}}
""")

    brand_param = 'brand := fs.String("brand", "", "sub-brand")\n\t' if brand_flag else ""
    brand_arg = "*brand" if brand_flag else '""'

    write(ROOT / f"{slug}-cli/cmd/{slug}/main.go", f"""// Command {slug} is an unofficial, agent-friendly CLI for {name}.
package main

import (
\t"fmt"
\t"os"
)

var version = "dev"

func main() {{
\tif len(os.Args) < 2 {{
\t\tusage()
\t\tos.Exit(2)
\t}}
\tvar err error
\tswitch os.Args[1] {{
\tcase "search":
\t\terr = cmdSearch(os.Args[2:])
\tcase "read":
\t\terr = cmdRead(os.Args[2:])
\tcase "brands":
\t\tcmdBrands()
\tcase "version", "--version", "-v":
\t\tfmt.Println(version)
\tcase "help", "-h", "--help":
\t\tusage()
\tdefault:
\t\tfmt.Fprintf(os.Stderr, "unknown command %q\\n\\n", os.Args[1])
\t\tusage()
\t\tos.Exit(2)
\t}}
\tif err != nil {{
\t\tfmt.Fprintln(os.Stderr, "error:", err)
\t\tos.Exit(1)
\t}}
}}

func usage() {{
\tfmt.Fprintf(os.Stderr, `{slug} — unofficial CLI for {name}

USAGE:
  {slug} search [--json]{' [--brand BRAND]' if brand_flag else ''} --from ORIGIN --to DEST --depart DATE [--return DATE]
  {slug} read [--json]{' [--brand BRAND]' if brand_flag else ''} <id|url>
  {slug} brands
  {slug} version | help
`)
}}
""")

    write(ROOT / f"{slug}-cli/cmd/{slug}/cli.go", """package main

import (
	"encoding/json"
	"flag"
	"os"
	"strings"
)

type common struct {
	jsonOut bool
}

func addCommon(fs *flag.FlagSet) *common {
	c := &common{}
	fs.BoolVar(&c.jsonOut, "json", false, "emit JSON to stdout")
	return c
}

func reorderArgs(fs *flag.FlagSet, args []string) []string {
	var flags, positional []string
	for i := 0; i < len(args); i++ {
		a := args[i]
		if a == "--" {
			positional = append(positional, args[i+1:]...)
			break
		}
		if len(a) > 1 && a[0] == '-' {
			flags = append(flags, a)
			name := strings.TrimLeft(a, "-")
			if strings.IndexByte(name, '=') >= 0 {
				continue
			}
			if f := fs.Lookup(name); f != nil {
				if _, ok := f.Value.(interface{ IsBoolFlag() bool }); !ok && i+1 < len(args) {
					flags = append(flags, args[i+1])
					i++
				}
			}
			continue
		}
		positional = append(positional, a)
	}
	out := append(flags, "--")
	return append(out, positional...)
}

func emitJSON(v any) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)
	return enc.Encode(v)
}
""")

    write(ROOT / f"{slug}-cli/cmd/{slug}/search.go", f"""package main

import (
\t"flag"
\t"fmt"

\t"github.com/fbelchi/{slug}-cli/internal/client"
)

func cmdSearch(args []string) error {{
\tfs := flag.NewFlagSet("search", flag.ExitOnError)
\tcf := addCommon(fs)
\t{brand_param}from := fs.String("from", "", "origin airport/city (IATA)")
\tto := fs.String("to", "", "destination airport/city (IATA)")
\tdepart := fs.String("depart", "", "departure date (YYYY-MM-DD)")
\tret := fs.String("return", "", "return date (YYYY-MM-DD, optional)")
\tlimit := fs.Int("limit", 24, "max results")
\tpage := fs.Int("page", 1, "page number")
\t_ = fs.Parse(reorderArgs(fs, args))
\tif *from == "" || *to == "" || *depart == "" {{
\t\treturn fmt.Errorf("usage: {slug} search [flags] (requires --from --to --depart)")
\t}}
\tcl := client.New({brand_arg})
\tres, err := cl.Search(*from, *to, *depart, *ret, *page, *limit)
\tif err != nil {{
\t\treturn err
\t}}
\tif cf.jsonOut {{
\t\treturn emitJSON(res)
\t}}
\tfmt.Printf("%s→%s depart=%s total=%d\\n", res.Origin, res.Dest, res.Depart, res.Total)
\tfor _, f := range res.Flights {{
\t\tfmt.Printf("  [%s] %s %s→%s %s — %s\\n", f.ID, f.FlightNumber, f.Origin, f.Destination, f.Depart, f.Price)
\t}}
\treturn nil
}}
""")

    write(ROOT / f"{slug}-cli/cmd/{slug}/read.go", f"""package main

import (
\t"flag"
\t"fmt"

\t"github.com/fbelchi/{slug}-cli/internal/client"
)

func cmdRead(args []string) error {{
\tfs := flag.NewFlagSet("read", flag.ExitOnError)
\tcf := addCommon(fs)
\t{brand_param}_ = fs.Parse(reorderArgs(fs, args))
\tif fs.NArg() != 1 {{
\t\treturn fmt.Errorf("usage: {slug} read [flags] <id|url>")
\t}}
\tcl := client.New({brand_arg})
\tfv, err := cl.Read(fs.Arg(0))
\tif err != nil {{
\t\treturn err
\t}}
\tif cf.jsonOut {{
\t\treturn emitJSON(fv)
\t}}
\tfmt.Printf("[%s] %s %s→%s %s\\n", fv.ID, fv.FlightNumber, fv.Origin, fv.Destination, fv.Depart)
\tif fv.Price.Price != "" {{
\t\tfmt.Printf("  price: %s %s\\n", fv.Price.Price, fv.Price.Currency)
\t}}
\treturn nil
}}
""")

    write(ROOT / f"{slug}-cli/cmd/{slug}/brands.go", f"""package main

import (
\t"fmt"

\t"github.com/fbelchi/{slug}-cli/internal/client"
)

func cmdBrands() {{
\tfor _, b := range client.Brands {{
\t\tfmt.Println(b)
\t}}
}}
""")

    brands_md = "\n".join(f"- {b}" for b in g["brands"])
    brand_section = ""
    if brand_flag:
        brand_section = f"""
## Sub-brands

{brands_md}

Use `--brand` to select a sub-brand.
"""

    write(ROOT / f"{slug}-cli/README.md", f"""# {name} CLI

Unofficial, agent-friendly CLI for [{name}]({g['url']}).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o {slug} ./cmd/{slug}
```

## Commands

```bash
{slug} search [--json]{' [--brand BRAND]' if brand_flag else ''} --from MAD --to BCN --depart 2026-07-01
{slug} read [--json]{' [--brand BRAND]' if brand_flag else ''} <id|url>
{slug} brands
```

## Environment

- `{env_prefix(slug)}_COOKIE` — optional browser cookie when blocked
- `{env_prefix(slug)}_REQUEST_DELAY` — rate limit (e.g. `2s`)
{brand_section}
## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
""")


def gen_readme(groups: list[dict]) -> None:
    hotels = [g for g in groups if g["cat"] == "hotel"]
    airlines = [g for g in groups if g["cat"] == "airline"]
    total_brands = sum(len(g["brands"]) for g in groups)

    def table_rows(gs):
        lines = []
        for g in gs:
            brands = ", ".join(g["brands"][:3])
            if len(g["brands"]) > 3:
                brands += f" (+{len(g['brands'])-3} more)"
            lines.append(f"| {g['name']} | [`{g['slug']}-cli/`]({g['slug']}-cli/) | `{g['slug']}` | {brands} | [README]({g['slug']}-cli/README.md) |")
        return "\n".join(lines)

    content = f"""# agentic-travel

Monorepo de **CLIs agent-friendly** para hoteles y aerolíneas (cadenas españolas e internacionales). Cada proyecto es un binario Go estático con salida `--json`, pensado para orquestación por agentes de IA.

> **No oficial.** APIs reverse-engineered. Ejecutar **solo en local** (IP residencial). Respeta rate limits.

## Resumen

- **{len(hotels)}** CLIs de hoteles, **{len(airlines)}** CLIs de aerolíneas
- **{total_brands}** marcas cubiertas (agrupadas por API padre compartida)
- Librería compartida: [`travelkit/`](travelkit/)

## Hoteles

| Grupo / API | Directorio | Binario | Marcas | README |
|-------------|------------|---------|--------|--------|
{table_rows(hotels)}

## Aerolíneas

| Grupo / API | Directorio | Binario | Marcas | README |
|-------------|------------|---------|--------|--------|
{table_rows(airlines)}

## Agrupación de marcas

Varias marcas comparten la misma API de reservas del grupo matriz. En esos casos hay **un solo CLI** con flag `--brand` y subcomando `brands`. Ejemplos:

| CLI | Marcas agrupadas |
|-----|------------------|
| `melia` | Meliá, Gran Meliá, Paradisus, INNSiDE, Sol, ZEL, … |
| `marriott` | Marriott, Ritz-Carlton, W Hotels, Sheraton, AC Hotels, … |
| `accor` | Ibis, Novotel, Mercure, Sofitel, Fairmont, … |
| `lufthansagroup` | Lufthansa, Swiss, Austrian, Eurowings, … |
| `turkish` | Turkish Airlines, Pegasus, SunExpress, AJet |

Ver `scripts/scaffold-clis.py` para el mapa completo grupo → marcas.

## Documentación para agentes

- **[AGENTS.md](AGENTS.md)** — guía general para agentes autónomos
- **[CLAUDE.md](CLAUDE.md)** — quick start para Cursor / Claude Code

## Requisitos

- Go 1.26+
- Chrome/Chromium headed browser for `session chrome --wait --timeout 3m` (Akamai `_abck` + `bm_sz`)

## Build rápido

```bash
cd melia-cli
go build -o melia ./cmd/melia
./melia search --json Madrid
```

## Verificación

```bash
./scripts/verify-clis.sh
```

## Estructura

```
agentic-travel/
├── travelkit/          # tipos, transport uTLS, cookies, rate limit
├── melia-cli/
├── ryanair-cli/
├── scripts/
│   ├── scaffold-clis.py
│   └── verify-clis.sh
└── …
```

Los CLIs declaran `replace github.com/fbelchi/travelkit => ../travelkit` en su `go.mod`.

## Licencia

Ver cada subproyecto. Uso bajo tu propia responsabilidad.
"""
    write(ROOT / "README.md", content)


def main() -> None:
    for g in GROUPS:
        if g["cat"] == "hotel":
            gen_hotel_cli(g)
        else:
            gen_airline_cli(g)
        print(f"scaffolded {g['slug']}-cli ({g['cat']})")
    gen_readme(GROUPS)
    manifest = [{"slug": g["slug"], "name": g["name"], "cat": g["cat"], "brands": g["brands"]} for g in GROUPS]
    write(ROOT / "scripts/groups.json", json.dumps(manifest, indent=2, ensure_ascii=False) + "\n")
    session_script = ROOT / "scripts/add-session-subcommands.py"
    if session_script.exists():
        subprocess.run([sys.executable, str(session_script)], check=True)
    print(f"Done: {len(GROUPS)} CLIs")


if __name__ == "__main__":
    main()
