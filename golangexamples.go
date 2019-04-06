package golangexamples
import "github.com/ehteshamz/greetings"

// returns the contents of the slice concatenated together and separated by a dash (-).
func ConcatSlice(sliceToConcat []byte) string {
  concatedString := ""
  for i,elem := range sliceToConcat{
    concatedString+=string(elem)
    if i!=len(sliceToConcat)-1 {
      concatedString+="-"
    }
  }
  return concatedString
}

func Encrypt(sliceToEncrypt []byte, ceaserCount int) {
  for i,_ := range sliceToEncrypt{
    sliceToEncrypt[i]+=byte(ceaserCount)
  }
}

func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}
