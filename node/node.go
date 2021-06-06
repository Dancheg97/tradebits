package node

type node struct {
	privKey     []byte
	pubKey      []byte
	ownerAdress []byte
	connections []byte
	unsynced    []byte
}

/*
# Распределение
Распределение будет производится по принципу 
*/


/*
# Транзакции
Добавить кадой транзакции поле, предполагающее подпись ноды, если данное поле
не несет нулевое значение, то транзакция считается проведенной на одной ноде и будет ждать в течении 1с. Ответа от первой ноды для проведения транзакции.
*/

func (n *node) Start() {

}

func (n *node) Connect() {

}

func (n *node) writeTransaction() {

}

func (n *node) 