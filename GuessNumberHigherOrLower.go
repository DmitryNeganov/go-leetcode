//374. Guess Number Higher or Lower
/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    left := 1
    right := n

    for left <= right {
        cur := (left + right) / 2
        guessing := guess(cur)
        if guessing == 1 {
            left = cur + 1
        } else if guessing == -1 {
            right = cur - 1
        } else {
            return cur
        }
    }
    
    return 0
}
