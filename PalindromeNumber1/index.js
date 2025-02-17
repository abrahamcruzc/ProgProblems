function isPalindrome(x) {
    var str = x.toString();
    var left = 0;
    var right = str.length - 1;
    while (left < right) {
        if (str[left] != str[right]) {
            return false;
        }
        left++;
        right--;
    }
    return true;
}
console.log(isPalindrome(12));
