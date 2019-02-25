package finident

import "strings"

var validCFI = map[string][]string{
	// Equities
	"ES": []string{"VNRE", "TU", "OPF", "BRNM"},
	"EP": []string{"VNRE", "RETGACN", "FCPQANU", "BRNM"},
	"EC": []string{"VNRE", "TU", "OPF", "BRNM"},
	"EF": []string{"VNRE", "RETGACN", "FCPQANU", "BRNM"},
	"EL": []string{"VNRE", "TU", "OPF", "BRNM"},
	"ED": []string{"SPCFLM", "RNBDX", "FCPQANUD", "BRNM}"},
	"EY": []string{"ABCDEM", "DYM", "FVEM", "BSDGTCINM"},
	"EM": []string{"X", "X", "X", "BRNM"},
	// Debt
	"DB": []string{"FZVCK", "TGSUPNOQJC", "FGCDABTLPQRE", "BRNM"},
	"DC": []string{"FZVK", "TGSUPNOQJC", "FGCDABTLPQRE", "BRNM"},
	"DW": []string{"FZVK", "TGSUPNOQJC", "FGCDABTLPQRE", "BRNM"},
	"DT": []string{"FZVK", "TGSUPNOQJC", "FGCDABTLPQRE", "BRNM"},
	"DY": []string{"FZVK", "TGSUPNOQJCX", "X", "BRNM"},
	"DS": []string{"ABCDM", "FDVYM", "FVM", "BSDTCINM"},
	"DE": []string{"ABCDEM", "FDVYM", "RSCTM", "BSDTCINM"},
	"DG": []string{"FZV", "TGSUPNOQJC", "FGCDABTLPQRE", "BRNM"},
	"DA": []string{"FZV", "TGSUPNOQJC", "FGCDABTLPQRE", "BRNM"},
	"DN": []string{"FZV", "TGSUPNOQJC", "FGCDABTLPQRE", "BRNM"},
	"DD": []string{"BCWTYGANM", "FZVC", "TGSUPNOQJC", "FGCDABTLPQRE"},
	"DM": []string{"BPM", "X", "X", "BRNM"},
	// Collective investment vehicles
	"CI": []string{"COM", "IGJ", "RBEVLCDFKM", "SQUY"},
	"CH": []string{"DRSEANLM", "X", "X", "X"},
	"CB": []string{"COM", "IGJ", "X", "SQUY"},
	"CE": []string{"COM", "IGJ", "RBEVLCDFKM", "SU"},
	"CS": []string{"COM", "BGLM", "RBM", "SU"},
	"CF": []string{"COM", "IGJ", "IHBEPM", "SQUY"},
	"CP": []string{"COM", "IGJ", "RBEVLCDFKM", "SQUY"},
	"CM": []string{"X", "X", "X", "SQUY"},
	// Entitlements
	"RA": []string{"X", "X", "X", "BRNM"},
	"RS": []string{"SPCFBIM", "X", "X", "BRNM"},
	"RP": []string{"SPCFBIM", "X", "X", "BRNM"},
	"RW": []string{"BSDTCIM", "TNC", "CPB", "EABM"},
	"RF": []string{"BSDTCIM", "TNM", "CPM", "EABM"},
	"RD": []string{"ASPWM", "X", "X", "BRNM"},
	"RM": []string{"X", "X", "X", "X"},
	// Listed options (OC & OP)
	"OC": []string{"EAB", "BSDTCIOFWNM", "PCNE", "SN"},
	"OP": []string{"EAB", "BSDTCIOFWNM", "PCNE", "SN"},
	"OM": []string{"X", "X", "X", "X"},
	// Futures
	"FF": []string{"BSDCIOFWNVM", "PCN", "SN", "X"},
	"FC": []string{"EAISNPHM", "PCN", "SN", "X"},
	// Swaps
	"SR": []string{"ACDGHZM", "CIDY", "SC", "CP"},
	"SF": []string{"ACM", "X", "X", "PN"},
	"ST": []string{"JKANGPSTIQM", "CT", "X", "CPE"},
	"SE": []string{"SIBM", "PDVLTCM", "X", "CPE"},
	"SC": []string{"UVIBM", "CTM", "CSL", "CPA"},
	"SM": []string{"PM", "X", "X", "CP"},
	// Non listed complex options
	"HR": []string{"ACDGHORFM", "ABCDEFGHI", "VADBGLPM", "CPE"},
	"HT": []string{"JKANGPSTIQORFWM", "ABCDEFGHI", "VADBGLPM", "CPE"},
	"HE": []string{"SIBORFM", "ABCDEFGHI", "VADBGLPM", "CPE"},
	"HC": []string{"UVIWM", "ABCDEFGHI", "VADBGLPM", "CPE"},
	"HF": []string{"RFTVM", "ABCDEFGHI", "VADBGLPM", "CPEN"},
	"HM": []string{"PM", "ABCDEFGHI", "VADBGLPM", "CPENA"},
	// Spot
	"IF": []string{"X", "X", "X", "P"},
	"IT": []string{"AJKNPSTM", "X", "X", "X"},
	// Forwards
	"JE": []string{"SIBOF", "X", "CSF", "CP"},
	"JF": []string{"TROF", "X", "CSF", "CPN"},
	"JC": []string{"AIBCDGO", "X", "SF", "CP"},
	"JR": []string{"IOM", "X", "SF", "CP"},
	"JT": []string{"ABGIJKNPSTM", "X", "CF", "CP"},
	// Strategies
	"KR": []string{"X", "X", "X", "X"},
	"KT": []string{"X", "X", "X", "X"},
	"KE": []string{"X", "X", "X", "X"},
	"KC": []string{"X", "X", "X", "X"},
	"KF": []string{"X", "X", "X", "X"},
	"KY": []string{"X", "X", "X", "X"},
	"KM": []string{"X", "X", "X", "X"},
	// Financing
	"LL": []string{"ABJKNPSTM", "X", "X", "CP"},
	"LR": []string{"GSC", "FNOT", "X", "DHT"},
	"LS": []string{"CGPTELDWKM", "NOT", "X", "DFHT"},
	// Referential instruments
	"TC": []string{"NLCM", "X", "X", "X"},
	"TT": []string{"EAISNPHM", "X", "X", "X"},
	"TR": []string{"NVFRM", "DWNQSAM", "X", "X"},
	"TI": []string{"EDFRTCM", "PCEFM", "PNGM", "X"},
	"TB": []string{"EDFITCM", "X", "X", "X"},
	"TD": []string{"SPCFLKM", "X", "X", "X"},
	"TM": []string{"X", "X", "X", "X"},
	// Others
	"MC": []string{"SBHAWUM", "TUX", "X", "BRNM"},
	"MM": []string{"RIETNPSM", "X", "X", "X"},
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
