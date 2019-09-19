package main

/**
空的select也会报错死锁的
*/
func main() {
	select {}
}
