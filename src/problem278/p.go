package problem278

import "sort"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 */
func isBadVersion(version int) bool {
	// implement me
	return true
}

func firstBadVersion(n int) int {
	idx := sort.Search(n, func(i int) bool { return isBadVersion(i + 1) })
	return idx + 1
}
