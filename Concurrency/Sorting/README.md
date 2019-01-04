This program sorts a sequence of integers provided by the user on the standard input in the following way:

* the input is parsed into an integer slice
* this slice is split into _numOfPieces_ number of equally or almost equally same length pieces
* each piece is sorted in an individual goroutine using the bubble sorting algorithm
* sync.WaitGroup is used for suspending the main thread until all goroutines finish their task
* when all sorting is finished, the pieces are merged into one sorted sequence

The implementation is an extension of an [earlier bubble sort algorithm project](https://github.com/uxexax/Go/tree/master/FunMethInt/BubbleSort "Bubble sort").