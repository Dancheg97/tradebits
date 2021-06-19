package node

import "crypto/rsa"

type node struct {
	privKey     *rsa.PrivateKey
	pubKey      []byte
	ownerAdress []byte
	presynced   []byte
	synced      []byte
}

/*
# Распределение
Распределение производится по следующим принципам:
 - Checker каждую минуту проверяет, осущестсвлялось ли распределение в текущий
 час, (проверяет базу данных со временем в виде байтов)
 -
*/

/*
# Транзакции
Добавить кадой транзакции поле, предполагающее подпись ноды, если данное поле
не несет нулевое значение, то транзакция считается проведенной на одной ноде и будет ждать в течении 1с. Ответа от первой ноды для проведения транзакции.
Каждая транзакция содержит следующий набор данных:
 - Состояние изменяемых объектов до измненения
 - Состояние изменяемых объектов после изменения
 - Время проведения транзакции
 - Хэш предыдущей транзакции (на которой текущая транзакция базируется)
*/

/*
# It is possible to connect to leading node by following algo:
 1) Getting connection request
 2) Checking minimal balance
 3) 

*/
func ConnectTo() {

}
