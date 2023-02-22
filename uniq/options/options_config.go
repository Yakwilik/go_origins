package options

const (
	eShowStrMeetCountFlag        = "c"
	showStrMeetCountFlagUsage    = "подсчитать количество вхождений строки во входных данных. Вывести это число перед строкой отделив пробелом."
	eShowNotUniqueStrFlag        = "d"
	showNotUniqueStrFlagUsage    = "вывести только те строки, которые повторились во входных данных."
	eShowUniqueStrFlag           = "u"
	showUniqueStrFlagUsage       = "вывести только те строки, которые не повторились во входных данных."
	skippedStringsCountFlag      = "f"
	skippedStringsCountFlagUsage = "не учитывать первые num_fields полей в строке. Полем в строке является непустой набор символов отделённый пробелом."
	skippedCharsCountFlag        = "s"
	skippedCharsCountFlagUsage   = "не учитывать первые num_chars символов в строке. При использовании вместе с параметром -f учитываются первые символы после num_fields полей (не учитывая пробел-разделитель после последнего поля)."
	ignoreRegisterFlag           = "i"
	ignoreRegisterUsage          = "не учитывать регистр букв."
)
