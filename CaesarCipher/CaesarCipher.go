package CaesarCipher

func Encode(message string, shift int) string {
  encoded := make([]byte, len(message))

  for _, runeValue := range message {

    selByte := ""

    if runeValue >= 97 && runeValue <= 122 {
      selByte = encode_char(runeValue, shift, 97)

    } else if runeValue >= 65 && runeValue <= 90 {
      selByte = encode_char(runeValue, shift, 65)

    } else {
      selByte = string(runeValue)

    }

    encoded = append(encoded, []byte(selByte)...)
  }

  return string(encoded[:])
}

func Decode(message string, shift int) string {
  return Encode(message, 26 - shift % 26)
}

func encode_char(code_point rune, shift int, ascii_diff int) string {
  encoded := ((int(code_point) - ascii_diff + shift) % 26 + ascii_diff)
  return string(encoded)
}
