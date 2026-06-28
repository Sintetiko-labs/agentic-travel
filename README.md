# agentic-travel

Monorepo de **CLIs agent-friendly** para hoteles y aerolíneas (cadenas españolas e internacionales). Cada proyecto es un binario Go estático con salida `--json`, pensado para orquestación por agentes de IA.

> **No oficial.** APIs reverse-engineered. Ejecutar **solo en local** (IP residencial). Respeta rate limits.

## Resumen

- **111** CLIs de hoteles, **83** CLIs de aerolíneas
- **321** marcas cubiertas (agrupadas por API padre compartida)
- Librería compartida: [`travelkit/`](travelkit/)

## Priority CLIs (loop 5)

| CLI | Tipo | Status | Session | Fuente / notas |
|-----|------|--------|---------|----------------|
| `ryanair` | airline | **live** | chrome+sync+doctor | `farfnd` + booking API |
| `vueling` | airline | **live** | chrome+sync+doctor | `apiwww.vueling.com/api/FlightPrice/GetAllFlights` (session optional) |
| `barcelo` | hotel | **live** | chrome+sync+doctor | JSON-LD listing |
| `riu` | hotel | **live** | chrome+sync+doctor | ng-state destination pages |
| `catalonia` | hotel | **live** | chrome+sync+doctor | homepage hotel links |
| `h10` | hotel | **live** | chrome+sync+doctor | ng-state `menu-es` destination pages |
| `palladium` | hotel | **live** | chrome+sync+doctor | AEM `data-hotel-name` cards |
| `lopesan` | hotel | **live** | chrome+sync+doctor | hotel detail links on listing pages |
| `princess` | hotel | **live** | chrome+sync+doctor | destination page headings |
| `melia` | hotel | partial | chrome+sync+doctor | BFF `/services/search/hotels/v2/search` (Akamai; needs `--wait`) |
| `nh` | hotel | partial | chrome+sync+doctor | REST `/nh/es/api/v1/hotels/search` (Akamai) |
| `iberostar` | hotel | partial | chrome+sync+doctor | GraphQL `/api/graphql` (Akamai) |
| `easyjet` | airline | partial | chrome+sync+doctor | ejavailability (Akamai after session) |
| `aireuropa` | airline | partial | chrome+sync+doctor | `dapi.aireuropa.com/api/v1/flights/search` POST |
| `iberiaexpress` | airline | partial | chrome+sync+doctor | `/api/availability/v1/flights` (Incapsula) |

**Session chrome (headed Chrome required):** `{slug} session chrome --wait --timeout 3m` polls until `_abck` **and** `bm_sz` (or `cf_clearance` / Incapsula pair). Saves to `~/.{slug}/cookies.json`. `{slug} session doctor` probes WAF cookies + brand API (POST bodies for BFF/GraphQL/dapi).

**CLIs with session subcommands:** **194** / 194 (via `scripts/add-session-subcommands.py`; scaffold regen runs it automatically).

**Smoke tests (loop 5):** `verify-clis.sh` → 194/194 PASS (build only). Live API smoke requires headed Chrome + residential IP for Akamai brands — not run in sandbox.

### Iteration 6 priorities

1. Priority partials — live search smoke after manual `session chrome --wait` (melia, nh, iberostar, easyjet, aireuropa, iberiaexpress)
2. Volotea / Binter — airline batch (simpler than Akamai-heavy majors)
3. Next hotel batch — eurostars, hotusa, vincci, silken, sercotel
4. Vueling — Skysales booking flow XHR (ScheduleSelect) for seat/fare detail beyond FlightPrice calendar
5. `session doctor` — extend POST probes to remaining partial CLIs

## Hoteles

| Grupo / API | Directorio | Binario | Marcas | README |
|-------------|------------|---------|--------|--------|
| Meliá | [`melia-cli/`](melia-cli/) | `melia` | Meliá Hotels International, Meliá, Gran Meliá (+6 more) | [README](melia-cli/README.md) |
| Barceló | [`barcelo-cli/`](barcelo-cli/) | `barcelo` | Barceló Hotel Group, Barceló Hotels & Resorts, Royal Hideaway (+2 more) | [README](barcelo-cli/README.md) |
| RIU | [`riu-cli/`](riu-cli/) | `riu` | RIU Hotels & Resorts | [README](riu-cli/README.md) |
| Iberostar | [`iberostar-cli/`](iberostar-cli/) | `iberostar` | Iberostar, Iberostar Selection, Iberostar Grand | [README](iberostar-cli/README.md) |
| NH Hotel Group | [`nh-cli/`](nh-cli/) | `nh` | NH Hotel Group, NH Hotels, NH Collection (+1 more) | [README](nh-cli/README.md) |
| Minor Hotels | [`minor-cli/`](minor-cli/) | `minor` | Avani, Tivoli, Minor Hotels | [README](minor-cli/README.md) |
| Eurostars | [`eurostars-cli/`](eurostars-cli/) | `eurostars` | Eurostars Hotel Company, Eurostars Hotels, Exe Hotels (+3 more) | [README](eurostars-cli/README.md) |
| Hotusa | [`hotusa-cli/`](hotusa-cli/) | `hotusa` | Hotusa, Crisol Hotels | [README](hotusa-cli/README.md) |
| H10 | [`h10-cli/`](h10-cli/) | `h10` | H10 Hotels, H10, Ocean by H10 | [README](h10-cli/README.md) |
| Princess Hotels | [`princess-cli/`](princess-cli/) | `princess` | Princess Hotels | [README](princess-cli/README.md) |
| Catalonia Hotels | [`catalonia-cli/`](catalonia-cli/) | `catalonia` | Catalonia Hotels & Resorts | [README](catalonia-cli/README.md) |
| Vincci | [`vincci-cli/`](vincci-cli/) | `vincci` | Vincci Hoteles | [README](vincci-cli/README.md) |
| Silken | [`silken-cli/`](silken-cli/) | `silken` | Silken Hoteles | [README](silken-cli/README.md) |
| Sercotel | [`sercotel-cli/`](sercotel-cli/) | `sercotel` | Sercotel | [README](sercotel-cli/README.md) |
| Room Mate | [`roommate-cli/`](roommate-cli/) | `roommate` | Room Mate Hotels | [README](roommate-cli/README.md) |
| Only YOU | [`onlyyou-cli/`](onlyyou-cli/) | `onlyyou` | Only YOU Hotels | [README](onlyyou-cli/README.md) |
| Palladium | [`palladium-cli/`](palladium-cli/) | `palladium` | Palladium Hotel Group, Ushuaïa Ibiza Beach Hotel, Hard Rock Hotel Ibiza (+3 more) | [README](palladium-cli/README.md) |
| BLESS Collection | [`bless-cli/`](bless-cli/) | `bless` | BLESS Collection Hotels | [README](bless-cli/README.md) |
| Grupo Piñero | [`pinero-cli/`](pinero-cli/) | `pinero` | Fiesta Hotels & Resorts, Grupo Piñero, Bahia Principe | [README](pinero-cli/README.md) |
| Senator | [`senator-cli/`](senator-cli/) | `senator` | Senator Hotels & Resorts, Playa Senator | [README](senator-cli/README.md) |
| Hipotels | [`hipotels-cli/`](hipotels-cli/) | `hipotels` | Hipotels | [README](hipotels-cli/README.md) |
| Lopesan | [`lopesan-cli/`](lopesan-cli/) | `lopesan` | Lopesan Hotel Group, Abora by Lopesan, Lopesan Hotels (+1 more) | [README](lopesan-cli/README.md) |
| Seaside Collection | [`seaside-cli/`](seaside-cli/) | `seaside` | Seaside Collection | [README](seaside-cli/README.md) |
| Be Live | [`belive-cli/`](belive-cli/) | `belive` | Be Live Hotels | [README](belive-cli/README.md) |
| Globales | [`globales-cli/`](globales-cli/) | `globales` | Globales Hotels, Hoteles Globales | [README](globales-cli/README.md) |
| Grupotel | [`grupotel-cli/`](grupotel-cli/) | `grupotel` | Grupotel | [README](grupotel-cli/README.md) |
| Garden Hotels | [`garden-cli/`](garden-cli/) | `garden` | Garden Hotels | [README](garden-cli/README.md) |
| Zafiro | [`zafiro-cli/`](zafiro-cli/) | `zafiro` | Zafiro Hotels | [README](zafiro-cli/README.md) |
| Viva Hotels | [`viva-cli/`](viva-cli/) | `viva` | Viva Hotels | [README](viva-cli/README.md) |
| Protur | [`protur-cli/`](protur-cli/) | `protur` | Protur Hotels | [README](protur-cli/README.md) |
| Fergus | [`fergus-cli/`](fergus-cli/) | `fergus` | Fergus Hotels | [README](fergus-cli/README.md) |
| Tent Hotels | [`tent-cli/`](tent-cli/) | `tent` | Tent Hotels | [README](tent-cli/README.md) |
| Iberik | [`iberik-cli/`](iberik-cli/) | `iberik` | Iberik Hoteles | [README](iberik-cli/README.md) |
| Servigroup | [`servigroup-cli/`](servigroup-cli/) | `servigroup` | Hoteles Servigroup | [README](servigroup-cli/README.md) |
| MedPlaya | [`medplaya-cli/`](medplaya-cli/) | `medplaya` | MedPlaya | [README](medplaya-cli/README.md) |
| Best Hotels | [`besthotels-cli/`](besthotels-cli/) | `besthotels` | Best Hotels | [README](besthotels-cli/README.md) |
| Alegria | [`alegria-cli/`](alegria-cli/) | `alegria` | Alegria Hotels | [README](alegria-cli/README.md) |
| HTop | [`htop-cli/`](htop-cli/) | `htop` | HTop Hotels | [README](htop-cli/README.md) |
| Guitart | [`guitart-cli/`](guitart-cli/) | `guitart` | Guitart Hotels | [README](guitart-cli/README.md) |
| Evenia | [`evenia-cli/`](evenia-cli/) | `evenia` | Evenia Hotels | [README](evenia-cli/README.md) |
| SB Hotels | [`sbhotels-cli/`](sbhotels-cli/) | `sbhotels` | SB Hotels | [README](sbhotels-cli/README.md) |
| Ilunion | [`ilunion-cli/`](ilunion-cli/) | `ilunion` | Ilunion Hotels | [README](ilunion-cli/README.md) |
| Paradores | [`paradores-cli/`](paradores-cli/) | `paradores` | Paradores | [README](paradores-cli/README.md) |
| Soho Boutique | [`soho-cli/`](soho-cli/) | `soho` | Soho Boutique Hotels | [README](soho-cli/README.md) |
| Casual Hoteles | [`casual-cli/`](casual-cli/) | `casual` | Casual Hoteles | [README](casual-cli/README.md) |
| Petit Palace | [`petitpalace-cli/`](petitpalace-cli/) | `petitpalace` | Petit Palace | [README](petitpalace-cli/README.md) |
| High Tech Hotels | [`hightech-cli/`](hightech-cli/) | `hightech` | High Tech Hotels | [README](hightech-cli/README.md) |
| One Shot | [`oneshot-cli/`](oneshot-cli/) | `oneshot` | One Shot Hotels | [README](oneshot-cli/README.md) |
| UMusic Hotels | [`umusic-cli/`](umusic-cli/) | `umusic` | UMusic Hotels | [README](umusic-cli/README.md) |
| Abba Hoteles | [`abba-cli/`](abba-cli/) | `abba` | Abba Hoteles | [README](abba-cli/README.md) |
| Zenit | [`zenit-cli/`](zenit-cli/) | `zenit` | Zenit Hoteles | [README](zenit-cli/README.md) |
| VP Hoteles | [`vp-cli/`](vp-cli/) | `vp` | VP Hoteles | [README](vp-cli/README.md) |
| Derby Hotels | [`derby-cli/`](derby-cli/) | `derby` | Derby Hotels Collection | [README](derby-cli/README.md) |
| Alma Hotels | [`alma-cli/`](alma-cli/) | `alma` | Alma Hotels | [README](alma-cli/README.md) |
| Hospes | [`hospes-cli/`](hospes-cli/) | `hospes` | Hospes Hotels | [README](hospes-cli/README.md) |
| Único Hotels | [`unico-cli/`](unico-cli/) | `unico` | Único Hotels | [README](unico-cli/README.md) |
| CoolRooms | [`coolrooms-cli/`](coolrooms-cli/) | `coolrooms` | CoolRooms Hotels | [README](coolrooms-cli/README.md) |
| Castilla Termal | [`castillatermal-cli/`](castillatermal-cli/) | `castillatermal` | Castilla Termal | [README](castillatermal-cli/README.md) |
| Eurobuilding | [`eurobuilding-cli/`](eurobuilding-cli/) | `eurobuilding` | Eurobuilding | [README](eurobuilding-cli/README.md) |
| Hoteles Center | [`center-cli/`](center-cli/) | `center` | Hoteles Center | [README](center-cli/README.md) |
| Hoteles Santos | [`santos-cli/`](santos-cli/) | `santos` | Hoteles Santos | [README](santos-cli/README.md) |
| Hoteles Elba | [`elba-cli/`](elba-cli/) | `elba` | Hoteles Elba | [README](elba-cli/README.md) |
| Hoteles Poseidón | [`poseidon-cli/`](poseidon-cli/) | `poseidon` | Hoteles Poseidón | [README](poseidon-cli/README.md) |
| Hoteles RH | [`rh-cli/`](rh-cli/) | `rh` | Hoteles RH | [README](rh-cli/README.md) |
| Magic Costa Blanca | [`magic-cli/`](magic-cli/) | `magic` | Magic Costa Blanca | [README](magic-cli/README.md) |
| Port Hotels | [`porthotels-cli/`](porthotels-cli/) | `porthotels` | Port Hotels | [README](porthotels-cli/README.md) |
| Ona Hotels | [`ona-cli/`](ona-cli/) | `ona` | Ona Hotels, Ona Hotels & Apartments | [README](ona-cli/README.md) |
| Pierre & Vacances | [`pierrevacances-cli/`](pierrevacances-cli/) | `pierrevacances` | Pierre & Vacances, Pierre & Vacances España, Apartamentos Pierre & Vacances | [README](pierrevacances-cli/README.md) |
| SmartRental | [`smartrental-cli/`](smartrental-cli/) | `smartrental` | SmartRental | [README](smartrental-cli/README.md) |
| Líbere | [`libere-cli/`](libere-cli/) | `libere` | Líbere Hospitality | [README](libere-cli/README.md) |
| Numa | [`numa-cli/`](numa-cli/) | `numa` | Numa | [README](numa-cli/README.md) |
| Sonder | [`sonder-cli/`](sonder-cli/) | `sonder` | Sonder | [README](sonder-cli/README.md) |
| Limehome | [`limehome-cli/`](limehome-cli/) | `limehome` | Limehome | [README](limehome-cli/README.md) |
| B&B Hotels | [`bbhotels-cli/`](bbhotels-cli/) | `bbhotels` | B&B Hotels | [README](bbhotels-cli/README.md) |
| Travelodge | [`travelodge-cli/`](travelodge-cli/) | `travelodge` | Travelodge | [README](travelodge-cli/README.md) |
| easyHotel | [`easyhotel-cli/`](easyhotel-cli/) | `easyhotel` | easyHotel | [README](easyhotel-cli/README.md) |
| Accor | [`accor-cli/`](accor-cli/) | `accor` | Ibis, Ibis Budget, Ibis Styles (+8 more) | [README](accor-cli/README.md) |
| Marriott | [`marriott-cli/`](marriott-cli/) | `marriott` | Marriott, Marriott Hotels, JW Marriott (+17 more) | [README](marriott-cli/README.md) |
| Hilton | [`hilton-cli/`](hilton-cli/) | `hilton` | Hilton, Hilton Hotels & Resorts, Conrad (+5 more) | [README](hilton-cli/README.md) |
| Hyatt | [`hyatt-cli/`](hyatt-cli/) | `hyatt` | Hyatt, Grand Hyatt, Hyatt Regency (+8 more) | [README](hyatt-cli/README.md) |
| IHG | [`ihg-cli/`](ihg-cli/) | `ihg` | IHG Hotels & Resorts, InterContinental, Kimpton (+6 more) | [README](ihg-cli/README.md) |
| Radisson | [`radisson-cli/`](radisson-cli/) | `radisson` | Radisson Hotel Group, Radisson Blu, Radisson RED (+2 more) | [README](radisson-cli/README.md) |
| Leonardo Hotels | [`leonardo-cli/`](leonardo-cli/) | `leonardo` | Leonardo Hotels, NYX Hotels, Leonardo Royal (+1 more) | [README](leonardo-cli/README.md) |
| Wyndham | [`wyndham-cli/`](wyndham-cli/) | `wyndham` | Wyndham Hotels & Resorts, Ramada, Wyndham (+2 more) | [README](wyndham-cli/README.md) |
| Best Western | [`bestwestern-cli/`](bestwestern-cli/) | `bestwestern` | Best Western, Best Western Plus, Best Western Premier (+1 more) | [README](bestwestern-cli/README.md) |
| Preferred Hotels | [`preferred-cli/`](preferred-cli/) | `preferred` | Preferred Hotels & Resorts | [README](preferred-cli/README.md) |
| Leading Hotels | [`lhw-cli/`](lhw-cli/) | `lhw` | Leading Hotels of the World | [README](lhw-cli/README.md) |
| Small Luxury Hotels | [`slh-cli/`](slh-cli/) | `slh` | Small Luxury Hotels of the World | [README](slh-cli/README.md) |
| Relais & Châteaux | [`relaischateaux-cli/`](relaischateaux-cli/) | `relaischateaux` | Relais & Châteaux | [README](relaischateaux-cli/README.md) |
| Design Hotels | [`designhotels-cli/`](designhotels-cli/) | `designhotels` | Design Hotels | [README](designhotels-cli/README.md) |
| Mandarin Oriental | [`mandarin-cli/`](mandarin-cli/) | `mandarin` | Mandarin Oriental | [README](mandarin-cli/README.md) |
| Four Seasons | [`fourseasons-cli/`](fourseasons-cli/) | `fourseasons` | Four Seasons | [README](fourseasons-cli/README.md) |
| Rosewood | [`rosewood-cli/`](rosewood-cli/) | `rosewood` | Rosewood | [README](rosewood-cli/README.md) |
| Belmond | [`belmond-cli/`](belmond-cli/) | `belmond` | Belmond | [README](belmond-cli/README.md) |
| Aman | [`aman-cli/`](aman-cli/) | `aman` | Aman | [README](aman-cli/README.md) |
| Nobu Hotels | [`nobu-cli/`](nobu-cli/) | `nobu` | Nobu Hotels | [README](nobu-cli/README.md) |
| Virgin Hotels | [`virgin-cli/`](virgin-cli/) | `virgin` | Virgin Hotels | [README](virgin-cli/README.md) |
| citizenM | [`citizenm-cli/`](citizenm-cli/) | `citizenm` | citizenM | [README](citizenm-cli/README.md) |
| Mama Shelter | [`mamashelter-cli/`](mamashelter-cli/) | `mamashelter` | Mama Shelter | [README](mamashelter-cli/README.md) |
| The Hoxton | [`hoxton-cli/`](hoxton-cli/) | `hoxton` | The Hoxton | [README](hoxton-cli/README.md) |
| 25hours Hotels | [`25hours-cli/`](25hours-cli/) | `25hours` | 25hours Hotels | [README](25hours-cli/README.md) |
| Ruby Hotels | [`ruby-cli/`](ruby-cli/) | `ruby` | Ruby Hotels | [README](ruby-cli/README.md) |
| Zoku | [`zoku-cli/`](zoku-cli/) | `zoku` | Zoku | [README](zoku-cli/README.md) |
| Locke | [`locke-cli/`](locke-cli/) | `locke` | Locke | [README](locke-cli/README.md) |
| ByPillow | [`bypillow-cli/`](bypillow-cli/) | `bypillow` | ByPillow | [README](bypillow-cli/README.md) |
| Axel Hotels | [`axel-cli/`](axel-cli/) | `axel` | Axel Hotels | [README](axel-cli/README.md) |
| Generator Hostels | [`generator-cli/`](generator-cli/) | `generator` | Generator Hostels | [README](generator-cli/README.md) |
| TOC Hostels | [`toc-cli/`](toc-cli/) | `toc` | TOC Hostels | [README](toc-cli/README.md) |
| Latroupe | [`latroupe-cli/`](latroupe-cli/) | `latroupe` | Latroupe | [README](latroupe-cli/README.md) |
| Safestay | [`safestay-cli/`](safestay-cli/) | `safestay` | Safestay | [README](safestay-cli/README.md) |
| St Christopher's | [`stchristophers-cli/`](stchristophers-cli/) | `stchristophers` | St Christopher's Inns Iberia | [README](stchristophers-cli/README.md) |

## Aerolíneas

| Grupo / API | Directorio | Binario | Marcas | README |
|-------------|------------|---------|--------|--------|
| Iberia Express | [`iberiaexpress-cli/`](iberiaexpress-cli/) | `iberiaexpress` | Iberia Express | [README](iberiaexpress-cli/README.md) |
| Vueling | [`vueling-cli/`](vueling-cli/) | `vueling` | Vueling | [README](vueling-cli/README.md) |
| Air Europa | [`aireuropa-cli/`](aireuropa-cli/) | `aireuropa` | Air Europa | [README](aireuropa-cli/README.md) |
| Ryanair | [`ryanair-cli/`](ryanair-cli/) | `ryanair` | Ryanair | [README](ryanair-cli/README.md) |
| easyJet | [`easyjet-cli/`](easyjet-cli/) | `easyjet` | easyJet | [README](easyjet-cli/README.md) |
| Wizz Air | [`wizzair-cli/`](wizzair-cli/) | `wizzair` | Wizz Air | [README](wizzair-cli/README.md) |
| Volotea | [`volotea-cli/`](volotea-cli/) | `volotea` | Volotea | [README](volotea-cli/README.md) |
| Binter | [`binter-cli/`](binter-cli/) | `binter` | Binter | [README](binter-cli/README.md) |
| Canaryfly | [`canaryfly-cli/`](canaryfly-cli/) | `canaryfly` | Canaryfly | [README](canaryfly-cli/README.md) |
| Level | [`level-cli/`](level-cli/) | `level` | Level | [README](level-cli/README.md) |
| Plus Ultra | [`plusultra-cli/`](plusultra-cli/) | `plusultra` | Plus Ultra Líneas Aéreas | [README](plusultra-cli/README.md) |
| World2Fly | [`world2fly-cli/`](world2fly-cli/) | `world2fly` | World2Fly | [README](world2fly-cli/README.md) |
| Iberojet | [`iberojet-cli/`](iberojet-cli/) | `iberojet` | Iberojet | [README](iberojet-cli/README.md) |
| Privilege Style | [`privilegestyle-cli/`](privilegestyle-cli/) | `privilegestyle` | Privilege Style | [README](privilegestyle-cli/README.md) |
| Air Nostrum | [`airnostrum-cli/`](airnostrum-cli/) | `airnostrum` | Air Nostrum | [README](airnostrum-cli/README.md) |
| Swiftair | [`swiftair-cli/`](swiftair-cli/) | `swiftair` | Swiftair | [README](swiftair-cli/README.md) |
| Albastar | [`albastar-cli/`](albastar-cli/) | `albastar` | Albastar | [README](albastar-cli/README.md) |
| TUI | [`tui-cli/`](tui-cli/) | `tui` | TUI Airways, TUI fly | [README](tui-cli/README.md) |
| Jet2 | [`jet2-cli/`](jet2-cli/) | `jet2` | Jet2 | [README](jet2-cli/README.md) |
| Norwegian | [`norwegian-cli/`](norwegian-cli/) | `norwegian` | Norwegian | [README](norwegian-cli/README.md) |
| Lufthansa Group | [`lufthansagroup-cli/`](lufthansagroup-cli/) | `lufthansagroup` | Lufthansa, Lufthansa City Airlines, Discover Airlines (+4 more) | [README](lufthansagroup-cli/README.md) |
| Air France-KLM | [`airfranceklm-cli/`](airfranceklm-cli/) | `airfranceklm` | Air France, KLM, Transavia | [README](airfranceklm-cli/README.md) |
| British Airways | [`britishairways-cli/`](britishairways-cli/) | `britishairways` | British Airways, BA CityFlyer | [README](britishairways-cli/README.md) |
| Aer Lingus | [`aerlingus-cli/`](aerlingus-cli/) | `aerlingus` | Aer Lingus | [README](aerlingus-cli/README.md) |
| TAP Air Portugal | [`tap-cli/`](tap-cli/) | `tap` | TAP Air Portugal | [README](tap-cli/README.md) |
| ITA Airways | [`ita-cli/`](ita-cli/) | `ita` | ITA Airways | [README](ita-cli/README.md) |
| SAS | [`sas-cli/`](sas-cli/) | `sas` | SAS | [README](sas-cli/README.md) |
| Finnair | [`finnair-cli/`](finnair-cli/) | `finnair` | Finnair | [README](finnair-cli/README.md) |
| LOT Polish Airlines | [`lot-cli/`](lot-cli/) | `lot` | LOT Polish Airlines | [README](lot-cli/README.md) |
| Czech Airlines | [`czechairlines-cli/`](czechairlines-cli/) | `czechairlines` | Czech Airlines | [README](czechairlines-cli/README.md) |
| Smartwings | [`smartwings-cli/`](smartwings-cli/) | `smartwings` | Smartwings | [README](smartwings-cli/README.md) |
| Aegean | [`aegean-cli/`](aegean-cli/) | `aegean` | Aegean Airlines, Olympic Air | [README](aegean-cli/README.md) |
| Croatia Airlines | [`croatiaairlines-cli/`](croatiaairlines-cli/) | `croatiaairlines` | Croatia Airlines | [README](croatiaairlines-cli/README.md) |
| Air Serbia | [`airserbia-cli/`](airserbia-cli/) | `airserbia` | Air Serbia | [README](airserbia-cli/README.md) |
| Turkish Airlines Group | [`turkish-cli/`](turkish-cli/) | `turkish` | Turkish Airlines, Pegasus Airlines, SunExpress (+1 more) | [README](turkish-cli/README.md) |
| Emirates | [`emirates-cli/`](emirates-cli/) | `emirates` | Emirates | [README](emirates-cli/README.md) |
| Qatar Airways | [`qatar-cli/`](qatar-cli/) | `qatar` | Qatar Airways | [README](qatar-cli/README.md) |
| Etihad Airways | [`etihad-cli/`](etihad-cli/) | `etihad` | Etihad Airways | [README](etihad-cli/README.md) |
| Saudia | [`saudia-cli/`](saudia-cli/) | `saudia` | Saudia | [README](saudia-cli/README.md) |
| Royal Jordanian | [`royaljordanian-cli/`](royaljordanian-cli/) | `royaljordanian` | Royal Jordanian | [README](royaljordanian-cli/README.md) |
| Kuwait Airways | [`kuwaitairways-cli/`](kuwaitairways-cli/) | `kuwaitairways` | Kuwait Airways | [README](kuwaitairways-cli/README.md) |
| Gulf Air | [`gulfair-cli/`](gulfair-cli/) | `gulfair` | Gulf Air | [README](gulfair-cli/README.md) |
| Egyptair | [`egyptair-cli/`](egyptair-cli/) | `egyptair` | Egyptair | [README](egyptair-cli/README.md) |
| Royal Air Maroc | [`royalairmaroc-cli/`](royalairmaroc-cli/) | `royalairmaroc` | Royal Air Maroc | [README](royalairmaroc-cli/README.md) |
| Air Arabia | [`airarabia-cli/`](airarabia-cli/) | `airarabia` | Air Arabia, Air Arabia Maroc | [README](airarabia-cli/README.md) |
| Tunisair | [`tunisair-cli/`](tunisair-cli/) | `tunisair` | Tunisair | [README](tunisair-cli/README.md) |
| Nouvelair | [`nouvelair-cli/`](nouvelair-cli/) | `nouvelair` | Nouvelair | [README](nouvelair-cli/README.md) |
| Air Algérie | [`airalgerie-cli/`](airalgerie-cli/) | `airalgerie` | Air Algerie | [README](airalgerie-cli/README.md) |
| Cabo Verde Airlines | [`caboverde-cli/`](caboverde-cli/) | `caboverde` | Cabo Verde Airlines | [README](caboverde-cli/README.md) |
| Ethiopian Airlines | [`ethiopian-cli/`](ethiopian-cli/) | `ethiopian` | Ethiopian Airlines | [README](ethiopian-cli/README.md) |
| Kenya Airways | [`kenyaairways-cli/`](kenyaairways-cli/) | `kenyaairways` | Kenya Airways | [README](kenyaairways-cli/README.md) |
| United Airlines | [`united-cli/`](united-cli/) | `united` | United Airlines | [README](united-cli/README.md) |
| American Airlines | [`american-cli/`](american-cli/) | `american` | American Airlines | [README](american-cli/README.md) |
| Delta Air Lines | [`delta-cli/`](delta-cli/) | `delta` | Delta Air Lines | [README](delta-cli/README.md) |
| Air Canada | [`aircanada-cli/`](aircanada-cli/) | `aircanada` | Air Canada | [README](aircanada-cli/README.md) |
| Air Transat | [`airtransat-cli/`](airtransat-cli/) | `airtransat` | Air Transat | [README](airtransat-cli/README.md) |
| WestJet | [`westjet-cli/`](westjet-cli/) | `westjet` | WestJet | [README](westjet-cli/README.md) |
| Aeroméxico | [`aeromexico-cli/`](aeromexico-cli/) | `aeromexico` | Aeroméxico | [README](aeromexico-cli/README.md) |
| Avianca | [`avianca-cli/`](avianca-cli/) | `avianca` | Avianca | [README](avianca-cli/README.md) |
| LATAM | [`latam-cli/`](latam-cli/) | `latam` | LATAM Airlines, LATAM Brasil, LATAM Chile | [README](latam-cli/README.md) |
| Aerolíneas Argentinas | [`aerolineas-cli/`](aerolineas-cli/) | `aerolineas` | Aerolíneas Argentinas | [README](aerolineas-cli/README.md) |
| Copa Airlines | [`copa-cli/`](copa-cli/) | `copa` | Copa Airlines | [README](copa-cli/README.md) |
| Air China | [`airchina-cli/`](airchina-cli/) | `airchina` | Air China | [README](airchina-cli/README.md) |
| China Eastern | [`chinaeastern-cli/`](chinaeastern-cli/) | `chinaeastern` | China Eastern | [README](chinaeastern-cli/README.md) |
| China Southern | [`chinasouthern-cli/`](chinasouthern-cli/) | `chinasouthern` | China Southern | [README](chinasouthern-cli/README.md) |
| Hainan Airlines | [`hainan-cli/`](hainan-cli/) | `hainan` | Hainan Airlines | [README](hainan-cli/README.md) |
| Cathay Pacific | [`cathaypacific-cli/`](cathaypacific-cli/) | `cathaypacific` | Cathay Pacific | [README](cathaypacific-cli/README.md) |
| Korean Air | [`koreanair-cli/`](koreanair-cli/) | `koreanair` | Korean Air | [README](koreanair-cli/README.md) |
| Asiana Airlines | [`asiana-cli/`](asiana-cli/) | `asiana` | Asiana Airlines | [README](asiana-cli/README.md) |
| Singapore Airlines | [`singaporeairlines-cli/`](singaporeairlines-cli/) | `singaporeairlines` | Singapore Airlines | [README](singaporeairlines-cli/README.md) |
| Thai Airways | [`thaiairways-cli/`](thaiairways-cli/) | `thaiairways` | Thai Airways | [README](thaiairways-cli/README.md) |
| Vietnam Airlines | [`vietnamairlines-cli/`](vietnamairlines-cli/) | `vietnamairlines` | Vietnam Airlines | [README](vietnamairlines-cli/README.md) |
| Qantas | [`qantas-cli/`](qantas-cli/) | `qantas` | Qantas | [README](qantas-cli/README.md) |
| El Al | [`elal-cli/`](elal-cli/) | `elal` | El Al | [README](elal-cli/README.md) |
| Icelandair | [`icelandair-cli/`](icelandair-cli/) | `icelandair` | Icelandair | [README](icelandair-cli/README.md) |
| PLAY Airlines | [`play-cli/`](play-cli/) | `play` | PLAY Airlines | [README](play-cli/README.md) |
| Norse Atlantic Airways | [`norse-cli/`](norse-cli/) | `norse` | Norse Atlantic Airways | [README](norse-cli/README.md) |
| Condor | [`condor-cli/`](condor-cli/) | `condor` | Condor | [README](condor-cli/README.md) |
| Corendon Airlines | [`corendon-cli/`](corendon-cli/) | `corendon` | Corendon Airlines | [README](corendon-cli/README.md) |
| Freebird Airlines | [`freebird-cli/`](freebird-cli/) | `freebird` | Freebird Airlines | [README](freebird-cli/README.md) |
| Enter Air | [`enterair-cli/`](enterair-cli/) | `enterair` | Enter Air | [README](enterair-cli/README.md) |
| Neos | [`neos-cli/`](neos-cli/) | `neos` | Neos | [README](neos-cli/README.md) |
| Wamos Air | [`wamos-cli/`](wamos-cli/) | `wamos` | Wamos Air | [README](wamos-cli/README.md) |

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
- Chrome/Chromium (solo para futuros comandos `session chrome`)

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
