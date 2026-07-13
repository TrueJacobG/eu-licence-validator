package validator

import (
	"regexp"
	"strings"
)

// plVoivodeships contains all valid first-letter voivodeship codes for
// standard and vintage plates. The primary letters are:
//   B C D E F G K L N O P R S T W Z
// Secondary letters introduced to expand capacity:
//   A (mazowieckie), I (śląskie), J (małopolskie), M (wielkopolskie),
//   V (dolnośląskie), X (pomorskie), Y (podkarpackie)
// H is NOT a voivodeship letter (it prefixes service plates).
// Q is never used in Polish plates. U prefixes military plates.
var plVoivodeships = map[byte]bool{
	'A': true, 'B': true, 'C': true, 'D': true, 'E': true, 'F': true,
	'G': true, 'I': true, 'J': true, 'K': true, 'L': true, 'M': true,
	'N': true, 'O': true, 'P': true, 'R': true, 'S': true, 'T': true,
	'V': true, 'W': true, 'X': true, 'Y': true, 'Z': true,
}

// plPowiatCodes is the complete lookup of all assigned 2- and 3-letter
// powiat/city-with-powiat-rights codes, sourced from the Wikisource
// "Polskie tablice rejestracyjne" page. Includes multi-codes for entities
// with more than one identifier, unused-but-assigned codes, and former codes.
var plPowiatCodes = map[string]bool{
	// dolnośląskie
	"DJ": true, "DL": true, "DB": true, "DW": true, "DX": true, "VW": true,
	"DBL": true, "DDZ": true, "DGL": true, "DGR": true, "DJA": true, "DJE": true,
	"DKA": true, "DKL": true, "DLE": true, "DLB": true, "DLU": true, "DLW": true,
	"DMI": true, "DOL": true, "DOA": true, "DPL": true, "DST": true, "DSR": true,
	"DSW": true, "DTR": true, "DBA": true, "DWL": true, "DWR": true, "VWR": true,
	"DZA": true, "DZG": true, "DZL": true,

	// kujawsko-pomorskie
	"CB": true, "CG": true, "CT": true, "CW": true,
	"CAL": true, "CBR": true, "CBY": true, "CBC": true, "CCH": true, "CGD": true,
	"CGR": true, "CIN": true, "CLI": true, "CMG": true, "CNA": true, "CRA": true,
	"CRY": true, "CSE": true, "CSW": true, "CTR": true, "CTU": true, "CWA": true,
	"CWL": true, "CZN": true,

	// lubelskie
	"LB": true, "LC": true, "LU": true, "LZ": true,
	"LBI": true, "LBL": true, "LCH": true, "LHR": true, "LJA": true, "LKS": true,
	"LKR": true, "LLB": true, "LUB": true, "LLE": true, "LLU": true, "LOP": true,
	"LPA": true, "LPU": true, "LRA": true, "LRY": true, "LSW": true, "LTM": true,
	"LWL": true, "LZA": true,

	// lubuskie
	"FG": true, "FZ": true,
	"FGW": true, "FKR": true, "FMI": true, "FNW": true, "FSL": true, "FSD": true,
	"FSU": true, "FSW": true, "FWS": true, "FZI": true, "FZG": true, "FZA": true,

	// łódzkie
	"EL": true, "ED": true, "EP": true, "ES": true,
	"EBE": true, "EBR": true, "EKU": true, "ELA": true, "ELE": true, "ELC": true,
	"ELW": true, "EOP": true, "EPA": true, "EPJ": true, "EPI": true, "EPD": true,
	"ERA": true, "ERW": true, "ESI": true, "ESK": true, "ETM": true, "EWI": true,
	"EWE": true, "EZD": true, "EZG": true,

	// małopolskie
	"KR": true, "KK": true, "KN": true, "KT": true,
	"KBC": true, "KBA": true, "KBR": true, "KCH": true, "KDA": true, "KGR": true,
	"KRA": true, "KRK": true, "KLI": true, "KMI": true, "KMY": true, "KNS": true,
	"KNT": true, "KOL": true, "KOS": true, "KPR": true, "KSU": true, "KTA": true,
	"KTT": true, "KWA": true, "KWI": true,

	// mazowieckie — Warsaw districts
	"WB": true, "WA": true, "WD": true, "WE": true, "WU": true, "WF": true,
	"WH": true, "WI": true, "WJ": true, "WK": true, "WN": true, "WT": true,
	"WY": true, "WX": true, "WW": true,
	// other cities with powiat rights
	"WO": true, "WP": true, "WR": true, "WS": true,
	// powiaty
	"WBR": true, "WCI": true, "WG": true, "WGS": true, "WGM": true, "WGR": true,
	"WKZ": true, "WL": true, "WLI": true, "WLS": true, "WM": true, "WMA": true,
	"WML": true, "WND": true, "WOS": true, "WOR": true, "WOT": true, "WPI": true,
	"WPA": true, "WPW": true, "WPX": true, "WPL": true, "WPN": true, "WPR": true,
	"WPP": true, "WPS": true, "WPZ": true, "WPY": true, "WPU": true, "WRA": true,
	"WSI": true, "WSE": true, "WSC": true, "WSK": true, "WSZ": true, "WWE": true,
	"WWL": true, "WV": true, "WWY": true, "WZ": true, "WZW": true, "WZU": true,
	"WZY": true,

	// opolskie
	"OP": true,
	"OB": true, "OGL": true, "OK": true, "OKL": true, "OKR": true, "ONA": true,
	"ONY": true, "OOL": true, "OPO": true, "OPR": true, "OST": true,

	// podkarpackie
	"RK": true, "RP": true, "RZ": true, "RT": true,
	"RBI": true, "RBR": true, "RDE": true, "RJA": true, "RJS": true, "YJS": true,
	"RKL": true, "RKR": true, "YKR": true, "RLS": true, "RLE": true, "RLU": true,
	"RLA": true, "RMI": true, "RNI": true, "RPR": true, "RPZ": true, "RRS": true,
	"RZE": true, "RZZ": true, "RSA": true, "RST": true, "RSR": true, "RTA": true,

	// podlaskie
	"BI": true, "BL": true, "BS": true,
	"BAU": true, "BIA": true, "BBI": true, "BGR": true, "BHA": true, "BKL": true,
	"BLM": true, "BMN": true, "BSE": true, "BSI": true, "BSK": true, "BSU": true,
	"BWM": true, "BZA": true,

	// pomorskie
	"GD": true, "GA": true, "GS": true, "GSP": true,
	"GBY": true, "GCH": true, "GCZ": true, "GDA": true, "GKA": true, "GKY": true,
	"GKZ": true, "GKS": true, "GKW": true, "GLE": true, "GMB": true, "GND": true,
	"GPU": true, "GSL": true, "GST": true, "XST": true, "GSZ": true, "GTC": true,
	"GWE": true, "GWO": true,

	// śląskie
	"SB": true, "SY": true, "SH": true, "SC": true, "SD": true, "SG": true,
	"SJZ": true, "SJ": true, "SK": true, "SM": true, "SPI": true, "SL": true,
	"SRS": true, "SR": true, "SI": true, "SO": true, "SW": true, "ST": true,
	"SZ": true, "SZO": true,
	"SBE": true, "SBI": true, "IBI": true, "SBL": true, "STY": true, "SCI": true,
	"SCN": true, "ICI": true, "SCZ": true, "SGL": true, "SKL": true, "SLU": true,
	"SMI": true, "SMY": true, "SPS": true, "SRC": true, "SRB": true, "STA": true,
	"ITA": true, "SWD": true, "SWZ": true, "SZA": true, "SZY": true,

	// świętokrzyskie
	"TK": true,
	"TBU": true, "TJE": true, "TKA": true, "TKI": true, "TKN": true, "TOP": true,
	"TOS": true, "TPI": true, "TSA": true, "TSK": true, "TST": true, "TSZ": true,
	"TLW": true,

	// warmińsko-mazurskie
	"NE": true, "NO": true,
	"NBA": true, "NBR": true, "NDZ": true, "NEB": true, "NEL": true, "NGI": true,
	"NGO": true, "NOG": true, "NIL": true, "NKE": true, "NMR": true, "NNM": true,
	"NOE": true, "NOL": true, "NOS": true, "NPI": true, "NSZ": true, "NWE": true,
	"NLI": true, "NNI": true,

	// wielkopolskie
	"PK": true, "PA": true, "PN": true, "PKO": true, "PL": true, "PO": true,
	"PY": true, "MO": true,
	"PCH": true, "PCT": true, "PGN": true, "PGS": true, "PGO": true, "PJA": true,
	"PKA": true, "PKE": true, "PKL": true, "PKN": true, "PKS": true, "PKR": true,
	"PLE": true, "PMI": true, "PNT": true, "POB": true, "POS": true, "POT": true,
	"PP": true, "PPL": true, "PZ": true, "POZ": true, "PRA": true, "PSL": true,
	"PSZ": true, "PSE": true, "PSR": true, "PTU": true, "MTU": true, "PWA": true,
	"PWL": true, "PWR": true, "PZL": true,

	// zachodniopomorskie
	"ZK": true, "ZS": true, "ZZ": true, "ZSW": true,
	"ZBI": true, "ZCH": true, "ZDR": true, "ZGL": true, "ZGY": true, "ZGR": true,
	"ZKA": true, "ZKL": true, "ZKO": true, "ZLO": true, "ZMY": true, "ZPL": true,
	"ZPY": true, "ZSL": true, "ZST": true, "ZSZ": true, "ZSD": true, "ZWA": true,
}

// plForbiddenLetters are letters banned in the vehicle identifier part
// (wyróżnik pojazdu) because they resemble digits: B→8, D→0, I→1, O→0, Z→2.
// Does NOT apply to individual (custom), service (H-prefix), or military
// (U-prefix) plates.
const plForbiddenLetters = "BDIOZ"

// plTempVoivodeships maps letter+digit codes for temporary/individual plates.
// Valid: B0-B9, C0-C9, D0-D9, V0-V9, E0-E9, F0-F9, G0-G9, X0-X9, K0-K9,
// J0-J9, L0-L9, N0-N9, O0-O9, P0-P9, M0-M9, R0-R9, Y0-Y9, S0-S9, I0-I9,
// T0-T9, W0-W9, A0-A9, Z0-Z9.
func plTempVoivodeshipOK(letter byte, digit byte) bool {
	if !plVoivodeships[letter] {
		return false
	}
	if digit < '0' || digit > '9' {
		return false
	}
	return true
}

func plIsDigit(c byte) bool { return c >= '0' && c <= '9' }
func plIsLetter(c byte) bool { return c >= 'A' && c <= 'Z' }

func plHasForbidden(s string) bool {
	return strings.ContainsAny(s, plForbiddenLetters)
}

// --- Regex patterns for standard plates (space optional) ---

// We use a broad matcher then validate the sequence by format.
var plRe2 = regexp.MustCompile(`^([A-Z]{2})\s?([0-9A-Z]{4,5})$`)
var plRe3 = regexp.MustCompile(`^([A-Z]{3})\s?([0-9A-Z]{3,5})$`)
var plReVintage2 = regexp.MustCompile(`^([A-Z]{2})\s?([0-9A-Z]{3})$`)
var plReVintage3 = regexp.MustCompile(`^([A-Z]{3})\s?([0-9A-Z]{2})$`)
var plReReduced = regexp.MustCompile(`^([A-Z])\s?([0-9A-Z]{3,4})$`)
var plReTemp = regexp.MustCompile(`^([A-Z])([0-9])\s?([0-9A-Z]{4,5})$`)
var plReTempResearch = regexp.MustCompile(`^([A-Z])([0-9])\s?(\d{3})\sB$`)
var plReIndividual = regexp.MustCompile(`^([A-Z])([0-9])\s?([A-Z0-9]{3,5})$`)
var plReDiplomatic = regexp.MustCompile(`^W\s?(\d{6})$`)
var plReMilitary = regexp.MustCompile(`^U([A-Z])\s?(\d{4,5})(T?)$`)
var plReService = regexp.MustCompile(`^(H[A-Z][A-Z]?)\s?([0-9A-Z]{4,5})$`)
var plReProfessional = regexp.MustCompile(`^([A-Z])(\d{2})\s?(\d{2})P(\d{2}|[1-9][A-Z])$`)

// --- Zasób validators for standard car/trailer plates ---

// plSeq2Car validates the vehicle-identifier part for 2-letter powiat
// car/trailer plates (zasoby I-V). seq is 4 or 5 chars.
func plSeq2Car(seq string) bool {
	switch len(seq) {
	case 5:
		// I: DDDDD  |  II: DDDDL  |  III: DDDLL  |  IV: D L DDD  |  V: D LL DD
		if onlyDigits(seq) {
			return true // I: 5 digits
		}
		if plIsDigit(seq[0]) && plIsDigit(seq[1]) && plIsDigit(seq[2]) && plIsDigit(seq[3]) && plIsLetter(seq[4]) {
			return !plHasForbidden(string(seq[4])) // II: 4 digits + 1 letter
		}
		if plIsDigit(seq[0]) && plIsDigit(seq[1]) && plIsDigit(seq[2]) && plIsLetter(seq[3]) && plIsLetter(seq[4]) {
			return !plHasForbidden(seq[3:]) // III: 3 digits + 2 letters
		}
		if plIsDigit(seq[0]) && seq[0] != '0' && plIsLetter(seq[1]) && plIsDigit(seq[2]) && plIsDigit(seq[3]) && plIsDigit(seq[4]) {
			return !plHasForbidden(string(seq[1])) // IV: 1 digit(non-0) + 1 letter + 3 digits
		}
		if plIsDigit(seq[0]) && seq[0] != '0' && plIsLetter(seq[1]) && plIsLetter(seq[2]) && plIsDigit(seq[3]) && plIsDigit(seq[4]) {
			return !plHasForbidden(seq[1:3]) // V: 1 digit(non-0) + 2 letters + 2 digits
		}
	case 4:
		// No 4-char zasob for 2-letter car plates
		return false
	}
	return false
}

// plSeq3Car validates the vehicle-identifier part for 3-letter powiat
// car/trailer plates (zasoby I-IX). seq is 4 or 5 chars.
func plSeq3Car(seq string) bool {
	switch len(seq) {
	case 5:
		// VII: DDDDD  |  VIII: DDDDL  |  IX: DDDLL
		if onlyDigits(seq) {
			return true // VII
		}
		if plIsDigit(seq[0]) && plIsDigit(seq[1]) && plIsDigit(seq[2]) && plIsDigit(seq[3]) && plIsLetter(seq[4]) {
			return !plHasForbidden(string(seq[4])) // VIII
		}
		if plIsDigit(seq[0]) && plIsDigit(seq[1]) && plIsDigit(seq[2]) && plIsLetter(seq[3]) && plIsLetter(seq[4]) {
			return !plHasForbidden(seq[3:]) // IX
		}
		return false
	case 4:
		// I: L DDD  |  II: DDLL  |  III: D L DD  |  IV: DD L D  |  V: D LL D  |  VI: LL DD
		if plIsLetter(seq[0]) && plIsDigit(seq[1]) && plIsDigit(seq[2]) && plIsDigit(seq[3]) {
			return !plHasForbidden(string(seq[0])) // I: 1 letter + 3 digits
		}
		if plIsDigit(seq[0]) && plIsDigit(seq[1]) && plIsLetter(seq[2]) && plIsLetter(seq[3]) {
			return !plHasForbidden(seq[2:]) // II: 2 digits + 2 letters
		}
		if plIsDigit(seq[0]) && seq[0] != '0' && plIsLetter(seq[1]) && plIsDigit(seq[2]) && plIsDigit(seq[3]) {
			return !plHasForbidden(string(seq[1])) // III: 1 digit(non-0) + 1 letter + 2 digits
		}
		if plIsDigit(seq[0]) && plIsDigit(seq[1]) && plIsLetter(seq[2]) && plIsDigit(seq[3]) && seq[3] != '0' {
			return !plHasForbidden(string(seq[2])) // IV: 2 digits + 1 letter + 1 digit(non-0)
		}
		if plIsDigit(seq[0]) && seq[0] != '0' && plIsLetter(seq[1]) && plIsLetter(seq[2]) && plIsDigit(seq[3]) && seq[3] != '0' {
			return !plHasForbidden(seq[1:3]) // V: 1 digit(non-0) + 2 letters + 1 digit(non-0)
		}
		if plIsLetter(seq[0]) && plIsLetter(seq[1]) && plIsDigit(seq[2]) && plIsDigit(seq[3]) {
			return !plHasForbidden(seq[0:2]) // VI: 2 letters + 2 digits
		}
		return false
	}
	return false
}

// plSeq2Moto validates motorcycle/moped 2-letter powiat plates (zasoby I-VIII).
// seq is 4 or 5 chars.
func plSeq2Moto(seq string) bool {
	switch len(seq) {
	case 4:
		// I: DDDD  |  II: DDDL  |  III: DD L D  |  IV: D L DD  |  V: L DDD  |  VI: DDLL  |  VII: D L L D  |  VIII: LLDD
		if onlyDigits(seq) {
			return true // I
		}
		if plIsDigit(seq[0]) && plIsDigit(seq[1]) && plIsDigit(seq[2]) && plIsLetter(seq[3]) {
			return !plHasForbidden(string(seq[3])) // II: 3 digits + 1 letter
		}
		if plIsDigit(seq[0]) && plIsDigit(seq[1]) && plIsLetter(seq[2]) && plIsDigit(seq[3]) && seq[3] != '0' {
			return !plHasForbidden(string(seq[2])) // III: 2 digits + 1 letter + 1 digit(non-0)
		}
		if plIsDigit(seq[0]) && seq[0] != '0' && plIsLetter(seq[1]) && plIsDigit(seq[2]) && plIsDigit(seq[3]) {
			return !plHasForbidden(string(seq[1])) // IV: 1 digit(non-0) + 1 letter + 2 digits
		}
		if plIsLetter(seq[0]) && plIsDigit(seq[1]) && plIsDigit(seq[2]) && plIsDigit(seq[3]) {
			return !plHasForbidden(string(seq[0])) // V: 1 letter + 3 digits
		}
		if plIsDigit(seq[0]) && plIsDigit(seq[1]) && plIsLetter(seq[2]) && plIsLetter(seq[3]) {
			return !plHasForbidden(seq[2:]) // VI: 2 digits + 2 letters
		}
		if plIsDigit(seq[0]) && seq[0] != '0' && plIsLetter(seq[1]) && plIsLetter(seq[2]) && plIsDigit(seq[3]) && seq[3] != '0' {
			return !plHasForbidden(seq[1:3]) // VII: 1 digit(non-0) + 2 letters + 1 digit(non-0)
		}
		if plIsLetter(seq[0]) && plIsLetter(seq[1]) && plIsDigit(seq[2]) && plIsDigit(seq[3]) {
			return !plHasForbidden(seq[0:2]) // VIII: 2 letters + 2 digits
		}
		return false
	}
	return false
}

// plSeq3Moto validates motorcycle/moped 3-letter powiat plates.
// Uses zasoby I-VI (same as car) plus X (L DD L) and XI (L D LL, D≠0).
// seq is 3, 4, or 5 chars.
func plSeq3Moto(seq string) bool {
	// zasoby I-VI are the same as car plates
	if plSeq3Car(seq) {
		return true
	}
	// X: L DD L  (4 chars)
	if len(seq) == 4 {
		if plIsLetter(seq[0]) && plIsDigit(seq[1]) && plIsDigit(seq[2]) && plIsLetter(seq[3]) {
			return !plHasForbidden(string(seq[0]) + string(seq[3])) // X: 1 letter + 2 digits + 1 letter
		}
	}
	// XI: L D LL  (4 chars, D≠0)
	if len(seq) == 4 {
		if plIsLetter(seq[0]) && plIsDigit(seq[1]) && seq[1] != '0' && plIsLetter(seq[2]) && plIsLetter(seq[3]) {
			return !plHasForbidden(string(seq[0]) + seq[2:]) // XI: 1 letter + 1 digit(non-0) + 2 letters
		}
	}
	return false
}

// plSeqReduced validates reduced plates (zmniejszone) — 1 voivodeship letter
// + zasoby I-VII. seq is 3 or 4 chars.
func plSeqReduced(seq string) bool {
	switch len(seq) {
	case 3:
		// I: DDD  |  II: DDL  |  III: D L D  |  IV: LDD  |  V: D LL  |  VI: LLD  |  VII: L D L
		if onlyDigits(seq) {
			return true // I: 3 digits
		}
		if plIsDigit(seq[0]) && plIsDigit(seq[1]) && plIsLetter(seq[2]) {
			return !plHasForbidden(string(seq[2])) // II: 2 digits + 1 letter
		}
		if plIsDigit(seq[0]) && seq[0] != '0' && plIsLetter(seq[1]) && plIsDigit(seq[2]) && seq[2] != '0' {
			return !plHasForbidden(string(seq[1])) // III: 1 digit(non-0) + 1 letter + 1 digit(non-0)
		}
		if plIsLetter(seq[0]) && plIsDigit(seq[1]) && plIsDigit(seq[2]) {
			return !plHasForbidden(string(seq[0])) // IV: 1 letter + 2 digits
		}
		if plIsDigit(seq[0]) && seq[0] != '0' && plIsLetter(seq[1]) && plIsLetter(seq[2]) {
			return !plHasForbidden(seq[1:]) // V: 1 digit(non-0) + 2 letters
		}
		if plIsLetter(seq[0]) && plIsLetter(seq[1]) && plIsDigit(seq[2]) && seq[2] != '0' {
			return !plHasForbidden(seq[0:2]) // VI: 2 letters + 1 digit(non-0)
		}
		if plIsLetter(seq[0]) && plIsDigit(seq[1]) && seq[1] != '0' && plIsLetter(seq[2]) {
			return !plHasForbidden(string(seq[0]) + string(seq[2])) // VII: 1 letter + 1 digit(non-0) + 1 letter
		}
		return false
	}
	return false
}

// plSeqVintage2 validates vintage plates for 2-letter powiat codes.
// Formats: XX DD L  or  XX DDD
func plSeqVintage2(seq string) bool {
	if len(seq) == 3 && onlyDigits(seq) {
		return true // 3 digits
	}
	if len(seq) == 3 && plIsDigit(seq[0]) && plIsDigit(seq[1]) && plIsLetter(seq[2]) {
		return !plHasForbidden(string(seq[2])) // 2 digits + 1 letter
	}
	return false
}

// plSeqVintage3 validates vintage plates for 3-letter powiat codes.
// Formats: XXX D L (D≠0)  |  XXX DD  |  XXX L D (D≠0)
func plSeqVintage3(seq string) bool {
	if len(seq) == 2 && onlyDigits(seq) {
		return true // 2 digits
	}
	if len(seq) == 2 && plIsDigit(seq[0]) && seq[0] != '0' && plIsLetter(seq[1]) {
		return !plHasForbidden(string(seq[1])) // 1 digit(non-0) + 1 letter
	}
	if len(seq) == 2 && plIsLetter(seq[0]) && plIsDigit(seq[1]) && seq[1] != '0' {
		return !plHasForbidden(string(seq[0])) // 1 letter + 1 digit(non-0)
	}
	return false
}

// --- Plate type validators ---

func validatePLStandard(plate string) bool {
	// Try 2-letter powiat
	if m := plRe2.FindStringSubmatch(plate); m != nil {
		region := m[1]
		seq := m[2]
		if !plPowiatCodes[region] {
			return false
		}
		// Try car/trailer zasoby
		if len(seq) == 5 && plSeq2Car(seq) {
			return true
		}
		// Try motorcycle/moped zasoby
		if plSeq2Moto(seq) {
			return true
		}
		return false
	}
	// Try 3-letter powiat
	if m := plRe3.FindStringSubmatch(plate); m != nil {
		region := m[1]
		seq := m[2]
		if !plPowiatCodes[region] {
			return false
		}
		// Try car/trailer zasoby
		if plSeq3Car(seq) {
			return true
		}
		// Try motorcycle/moped zasoby
		if plSeq3Moto(seq) {
			return true
		}
		return false
	}
	return false
}

func validatePLReduced(plate string) bool {
	m := plReReduced.FindStringSubmatch(plate)
	if m == nil {
		return false
	}
	voiv := m[1][0]
	seq := m[2]
	if !plVoivodeships[voiv] {
		return false
	}
	return plSeqReduced(seq)
}

func validatePLTemporary(plate string) bool {
	// Research/testing plates: LD DDD B (B always separated by space)
	if m := plReTempResearch.FindStringSubmatch(plate); m != nil {
		letter := m[1][0]
		digit := m[2][0]
		if plTempVoivodeshipOK(letter, digit) {
			return true
		}
		return false
	}
	// Standard temporary: LD DDDD  or  LD DDDL
	if m := plReTemp.FindStringSubmatch(plate); m != nil {
		letter := m[1][0]
		digit := m[2][0]
		seq := m[3]
		if !plTempVoivodeshipOK(letter, digit) {
			return false
		}
		// I zasób: 4 digits
		if len(seq) == 4 && onlyDigits(seq) {
			return true
		}
		// II zasób: 3 digits + 1 letter (B only valid in research format with space)
		if len(seq) == 4 && plIsDigit(seq[0]) && plIsDigit(seq[1]) && plIsDigit(seq[2]) && plIsLetter(seq[3]) {
			return !plHasForbidden(string(seq[3]))
		}
		return false
	}
	return false
}

func validatePLIndividual(plate string) bool {
	m := plReIndividual.FindStringSubmatch(plate)
	if m == nil {
		return false
	}
	letter := m[1][0]
	digit := m[2][0]
	custom := m[3]

	if !plTempVoivodeshipOK(letter, digit) {
		return false
	}

	n := len(custom)
	if n < 3 || n > 5 {
		return false
	}

	// First char must be a letter
	if !plIsLetter(custom[0]) {
		return false
	}

	// Q is never allowed (even on custom plates)
	if strings.ContainsRune(custom, 'Q') {
		return false
	}

	// All letters come before digits (no intermixing).
	// Max 2 trailing digits.
	// Find where digits start (if any).
	digitStart := -1
	for i := 0; i < n; i++ {
		if plIsDigit(custom[i]) {
			digitStart = i
			break
		}
	}
	if digitStart == -1 {
		// All letters — valid (3-5 letters)
		return true
	}
	// Digits at the end: all chars from digitStart to end must be digits
	for i := digitStart; i < n; i++ {
		if !plIsDigit(custom[i]) {
			return false // letter after digit = intermixed
		}
	}
	// Max 2 trailing digits
	trailingDigits := n - digitStart
	if trailingDigits > 2 {
		return false
	}
	return true
}

func validatePLVintage(plate string) bool {
	// 2-letter powiat — seq is 3 chars (DD L or DDD)
	if m := plReVintage2.FindStringSubmatch(plate); m != nil {
		region := m[1]
		seq := m[2]
		if !plPowiatCodes[region] {
			return false
		}
		return plSeqVintage2(seq)
	}
	// 3-letter powiat — seq is 2 chars (D L, DD, or L D)
	if m := plReVintage3.FindStringSubmatch(plate); m != nil {
		region := m[1]
		seq := m[2]
		if !plPowiatCodes[region] {
			return false
		}
		return plSeqVintage3(seq)
	}
	return false
}

func validatePLDiplomatic(plate string) bool {
	m := plReDiplomatic.FindStringSubmatch(plate)
	if m == nil {
		return false
	}
	digits := m[1]
	if len(digits) != 6 {
		return false
	}
	// First 3 digits = country code (001-121)
	countryCode := digits[0:3]
	// Must be 001-121
	val := int(countryCode[0]-'0')*100 + int(countryCode[1]-'0')*10 + int(countryCode[2]-'0')
	if val < 1 || val > 121 {
		return false
	}
	// Last 3 digits = purpose code — accept any 3 digits (001-999)
	purpose := digits[3:6]
	if purpose == "000" {
		return false
	}
	return true
}

func validatePLMilitary(plate string) bool {
	m := plReMilitary.FindStringSubmatch(plate)
	if m == nil {
		return false
	}
	serviceLetter := m[1][0]
	num := m[2]
	suffix := m[3]

	// Service letters: A,B,C,D,E,G,I,J,K,L (not O or U)
	switch serviceLetter {
	case 'A', 'B', 'C', 'D', 'E', 'G', 'I', 'J', 'K', 'L':
	default:
		return false
	}

	if suffix == "T" {
		// Tracked: 4 digits + T
		return len(num) == 4
	}
	// Cars: 5 digits, motorcycles: 4 digits
	return len(num) == 5 || len(num) == 4
}

func validatePLService(plate string) bool {
	m := plReService.FindStringSubmatch(plate)
	if m == nil {
		return false
	}
	prefix := m[1] // H + 1 or 2 service letters
	seq := m[2]

	// Validate the service code letter (second char)
	serviceLetter := prefix[1]
	switch serviceLetter {
	case 'A', 'B', 'C', 'K', 'M', 'P', 'W', 'S':
	default:
		return false
	}

	// Both I and II zasób have 4-char seq regardless of prefix length.
	if len(seq) != 4 {
		return false
	}
	// II zasób: DD LL (2 digits + 2 letters) — forbidden letters don't apply
	if plIsDigit(seq[0]) && plIsDigit(seq[1]) && plIsLetter(seq[2]) && plIsLetter(seq[3]) {
		return true
	}
	// I zasób: L DDD (1 letter + 3 digits)
	if plIsLetter(seq[0]) && onlyDigits(seq[1:]) {
		return true
	}
	return false
}

func validatePLProfessional(plate string) bool {
	m := plReProfessional.FindStringSubmatch(plate)
	if m == nil {
		return false
	}
	letter := m[1][0]
	if !plVoivodeships[letter] {
		return false
	}
	// m[2] = 2-digit powiat code
	// m[3] = 2 digits (00-99)
	// P is literal in the regex
	// m[4] = I zasób: 2 digits (01-99) | II zasób: 1 digit(1-9) + 1 letter
	last := m[4]
	if len(last) == 2 {
		if plIsDigit(last[0]) && plIsDigit(last[1]) && last != "00" {
			return true // I zasób: 2 digits (01-99)
		}
		if plIsDigit(last[0]) && last[0] >= '1' && last[0] <= '9' && plIsLetter(last[1]) {
			return !plHasForbidden(string(last[1])) // II zasób: 1 digit + 1 letter (no forbidden)
		}
	}
	return false
}

// validatePL tries all Polish plate types and returns true if any match.
func validatePL(plate string) bool {
	return validatePLStandard(plate) ||
		validatePLReduced(plate) ||
		validatePLTemporary(plate) ||
		validatePLIndividual(plate) ||
		validatePLVintage(plate) ||
		validatePLDiplomatic(plate) ||
		validatePLMilitary(plate) ||
		validatePLService(plate) ||
		validatePLProfessional(plate)
}
