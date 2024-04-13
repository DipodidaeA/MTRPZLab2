package lab2

import (
	"strings"
	"unicode/utf8"
)

func ChengeText(textInput string) (string, string) {
	textNew := ""
	var textSize int
	preOpen := 0
	ttOpen := 0
	bOpen := 0
	iOpen := 0
	err := ""
	textInput = strings.ReplaceAll(textInput, "\n", " ")
	textSize = utf8.RuneCountInString(textInput)
	textNew += "<p>"
	if textInput != "" {
		textInput += "FORONEERROR"
		for i := 0; i < textSize; i++ {
			if err == "" {
				switch textInput[i] {
				case '`':
					if textInput[i+1] == '`' {
						if textInput[i+2] == '`' {
							i += 3
							preOpen = 1
							textNew += "<pre>"
							for i < textSize {
								if preOpen == 0 {
									break
								} else {
									switch textInput[i] {
									case '`':
										if textInput[i+1] == '`' && textInput[i+2] == '`' && textInput[i+3] != '`' {
											preOpen = 0
											i += 2
											textNew += "</pre>"
										} else {
											textNew += string(textInput[i])
											i++
										}
									default:
										textNew += string(textInput[i])
										i++
									}
								}
							}
						} else {
							textNew += string(textInput[i])
						}
					} else {
						if textInput[i+1] != ' ' && textInput[i+1] != '*' && textInput[i+1] != '_' {
							i++
							ttOpen = 1
							textNew += "<tt>"
							for i < textSize {
								if ttOpen == 0 {
									break
								} else {
									switch textInput[i] {
									case '`':
										if textInput[i-1] != ' ' {
											ttOpen = 0
											textNew += "</tt>"
										} else {
											ttOpen = 0
											err = "Error char `"
										}
									default:
										textNew += string(textInput[i])
										i++
									}
								}
							}
						} else {
							textNew += string(textInput[i])
						}
					}
				case '*':
					if textInput[i+1] == '*' {
						if textInput[i+2] == '*' {
							textNew += string(textInput[i])
						} else {
							if textInput[i+2] != ' ' && textInput[i+2] != '`' && textInput[i+2] != '_' {
								i += 2
								bOpen = 1
								textNew += "<b>"
								for i < textSize {
									if bOpen == 0 {
										break
									} else {
										switch textInput[i] {
										case '*':
											if textInput[i-1] != ' ' && textInput[i+1] == '*' {
												bOpen = 0
												textNew += "</b>"
												i++
											} else {
												bOpen = 0
												err = "Error char *"
											}
										default:
											textNew += string(textInput[i])
											i++
										}
									}
								}
							} else {
								textNew += string(textInput[i])
							}
						}
					} else {
						textNew += string(textInput[i])
					}
				case '_':
					if textInput[i+1] == '_' {
						if textInput[i+2] == '_' {
							textNew += string(textInput[i])
						} else {
							textNew += string(textInput[i])
						}
					} else {
						if textInput[i+1] != ' ' && textInput[i+1] != '*' && textInput[i+1] != '`' {
							i++
							iOpen = 1
							textNew += "<i>"
							for i < textSize {
								if iOpen == 0 {
									break
								} else {
									switch textInput[i] {
									case '_':
										if textInput[i-1] != ' ' {
											iOpen = 0
											textNew += "</i>"
										} else {
											iOpen = 0
											err = "Error char _"
										}
									default:
										textNew += string(textInput[i])
										i++
									}
								}
							}
						} else {
							textNew += string(textInput[i])
						}
					}
				case ' ':
					if textInput[i+1] == ' ' {
						if textInput[i+2] == ' ' {
							break
						} else {
							textNew += "</p><p>"
							i++
						}
					} else {
						textNew += string(textInput[i])
					}
				default:
					textNew += string(textInput[i])
				}
			} else {
				break
			}
		}

		if err == "" && preOpen == 0 && ttOpen == 0 && iOpen == 0 && bOpen == 0 {
			textNew += "</p>"
			textNew = strings.ReplaceAll(textNew, "FORONEERROR", "")
		}
	} else {
		err = "Відсутній текст у файлі"
	}
	return textNew, err
}
