package finident

import "strings"

var validCFI = map[string][]string{
	// Equities
	"ES": {"VNRE", "TU", "OPF", "BRNM"},
	"EP": {"VNRE", "RETGACN", "FCPQANU", "BRNM"},
	"EC": {"VNRE", "TU", "OPF", "BRNM"},
	"EF": {"VNRE", "RETGACN", "FCPQANU", "BRNM"},
	"EL": {"VNRE", "TU", "OPF", "BRNM"},
	"ED": {"SPCFLM", "RNBDX", "FCPQANUD", "BRNM"},
	"EY": {"ABCDEM", "DYM", "FVEM", "BSDGTCINM"},
	"EM": {"X", "X", "X", "BRNM"},
	// Debt
	"DB": {"FZVCK", "TGSUPNOQJC", "FGCDABTLPQRE", "BRNM"},
	"DC": {"FZVK", "TGSUPNOQJC", "FGCDABTLPQRE", "BRNM"},
	"DW": {"FZVK", "TGSUPNOQJC", "FGCDABTLPQRE", "BRNM"},
	"DT": {"FZVK", "TGSUPNOQJC", "FGCDABTLPQRE", "BRNM"},
	"DY": {"FZVK", "TGSUPNOQJCX", "X", "BRNM"},
	"DS": {"ABCDM", "FDVYM", "FVM", "BSDTCINM"},
	"DE": {"ABCDEM", "FDVYM", "RSCTM", "BSDTCINM"},
	"DG": {"FZV", "TGSUPNOQJC", "FGCDABTLPQRE", "BRNM"},
	"DA": {"FZV", "TGSUPNOQJC", "FGCDABTLPQRE", "BRNM"},
	"DN": {"FZV", "TGSUPNOQJC", "FGCDABTLPQRE", "BRNM"},
	"DD": {"BCWTYGANM", "FZVC", "TGSUPNOQJC", "FGCDABTLPQRE"},
	"DM": {"BPM", "X", "X", "BRNM"},
	// Collective investment vehicles
	"CI": {"COM", "IGJ", "RBEVLCDFKM", "SQUY"},
	"CH": {"DRSEANLM", "X", "X", "X"},
	"CB": {"COM", "IGJ", "X", "SQUY"},
	"CE": {"COM", "IGJ", "RBEVLCDFKM", "SU"},
	"CS": {"COM", "BGLM", "RBM", "SU"},
	"CF": {"COM", "IGJ", "IHBEPM", "SQUY"},
	"CP": {"COM", "IGJ", "RBEVLCDFKM", "SQUY"},
	"CM": {"X", "X", "X", "SQUY"},
	// Entitlements
	"RA": {"X", "X", "X", "BRNM"},
	"RS": {"SPCFBIM", "X", "X", "BRNM"},
	"RP": {"SPCFBIM", "X", "X", "BRNM"},
	"RW": {"BSDTCIM", "TNC", "CPB", "EABM"},
	"RF": {"BSDTCIM", "TNM", "CPM", "EABM"},
	"RD": {"ASPWM", "X", "X", "BRNM"},
	"RM": {"X", "X", "X", "X"},
	// Listed options (OC & OP)
	"OC": {"EAB", "BSDTCIOFWNM", "PCNE", "SN"},
	"OP": {"EAB", "BSDTCIOFWNM", "PCNE", "SN"},
	"OM": {"X", "X", "X", "X"},
	// Futures
	"FF": {"BSDCIOFWNVM", "PCN", "SN", "X"},
	"FC": {"EAISNPHM", "PCN", "SN", "X"},
	// Swaps
	"SR": {"ACDGHZM", "CIDY", "SC", "CP"},
	"SF": {"ACM", "X", "X", "PN"},
	"ST": {"JKANGPSTIQM", "CT", "X", "CPE"},
	"SE": {"SIBM", "PDVLTCM", "X", "CPE"},
	"SC": {"UVIBM", "CTM", "CSL", "CPA"},
	"SM": {"PM", "X", "X", "CP"},
	// Non listed complex options
	"HR": {"ACDGHORFM", "ABCDEFGHI", "VADBGLPM", "CPE"},
	"HT": {"JKANGPSTIQORFWM", "ABCDEFGHI", "VADBGLPM", "CPE"},
	"HE": {"SIBORFM", "ABCDEFGHI", "VADBGLPM", "CPE"},
	"HC": {"UVIWM", "ABCDEFGHI", "VADBGLPM", "CPE"},
	"HF": {"RFTVM", "ABCDEFGHI", "VADBGLPM", "CPEN"},
	"HM": {"PM", "ABCDEFGHI", "VADBGLPM", "CPENA"},
	// Spot
	"IF": {"X", "X", "X", "P"},
	"IT": {"AJKNPSTM", "X", "X", "X"},
	// Forwards
	"JE": {"SIBOF", "X", "CSF", "CP"},
	"JF": {"TROF", "X", "CSF", "CPN"},
	"JC": {"AIBCDGO", "X", "SF", "CP"},
	"JR": {"IOM", "X", "SF", "CP"},
	"JT": {"ABGIJKNPSTM", "X", "CF", "CP"},
	// Strategies
	"KR": {"X", "X", "X", "X"},
	"KT": {"X", "X", "X", "X"},
	"KE": {"X", "X", "X", "X"},
	"KC": {"X", "X", "X", "X"},
	"KF": {"X", "X", "X", "X"},
	"KY": {"X", "X", "X", "X"},
	"KM": {"X", "X", "X", "X"},
	// Financing
	"LL": {"ABJKNPSTM", "X", "X", "CP"},
	"LR": {"GSC", "FNOT", "X", "DHT"},
	"LS": {"CGPTELDWKM", "NOT", "X", "DFHT"},
	// Referential instruments
	"TC": {"NLCM", "X", "X", "X"},
	"TT": {"EAISNPHM", "X", "X", "X"},
	"TR": {"NVFRM", "DWNQSAM", "X", "X"},
	"TI": {"EDFRTCM", "PCEFM", "PNGM", "X"},
	"TB": {"EDFITCM", "X", "X", "X"},
	"TD": {"SPCFLKM", "X", "X", "X"},
	"TM": {"X", "X", "X", "X"},
	// Others
	"MC": {"SBHAWUM", "TUX", "X", "BRNM"},
	"MM": {"RIETNPSM", "X", "X", "X"},
}

// IsValidCFI returns true if the CFI string is a valid CFI code, and false otherwise.
// This validator is based on ESMAs CFI list published here https://www.esma.europa.eu/file/20301/download?token=6K3VKc5m
func IsValidCFI(cfi string) bool {
	if len(cfi) != 6 {
		return false
	}
	attributes, ok := validCFI[cfi[0:2]]
	if !ok {
		return false
	}
	for i := range cfi[2:] {
		if strings.IndexByte(attributes[i], cfi[i+2]) == -1 {
			return false
		}
	}
	return true
}

// GenCFICombinations returns a list of all valid CFIs.
// This validator is based on ESMAs CFI list published here https://www.esma.europa.eu/file/20301/download?token=6K3VKc5m
func GenCFICombinations() []string {
	cfis := make([]string, 0, 10000)
	for k, v := range validCFI {
		for _, v2 := range combinations(v) {
			cfis = append(cfis, k+v2)
		}
	}
	return cfis
}

func combinations(v []string) []string {
	combs := []string{}
	curvals := v[0]
	// get combinations of next segment
	var c []string
	if len(v) > 1 {
		c = combinations(v[1:])
	} else {
		return strings.Split(curvals, "")
	}
	for i := range curvals {
		for j := range c {
			combs = append(combs, string(curvals[i])+c[j])
		}
	}
	return combs
}
