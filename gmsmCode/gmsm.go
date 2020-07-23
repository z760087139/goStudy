package gmsmcode

import (
	"github.com/tjfoc/gmsm/sm2"
)

func main() {
	private, _ := sm2.GenerateKey()
	public := private.PublicKey
	public.Encrypt()

}
