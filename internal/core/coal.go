package coal

import (
	"io"

	"gitlab.com/gomidi/midi/v2/smf"
)

// Start function for the Coal algorithm
func Start(payload []uint8, datasetSong *io.Reader) float64 {

	songToCompare := ReadNotesOn(datasetSong)

	// Call comparison system
	matchingSequence := lcs(payload, songToCompare)

	// fmt.Println(matchingSequence)

	// Percentage of matching notes in the input song
	matchingRate := float64(len(matchingSequence)) / float64(len(payload))

	return matchingRate
}

// Read only Notes On
func ReadNotesOn(song *io.Reader) []uint8 {
	var notesOn []uint8

	trackReader := smf.ReadTracksFrom(*song)

	trackReader.Do(func(te smf.TrackEvent) {
		var ch, key, vel uint8

		// Get only notes on
		if te.Message.GetNoteOn(&ch, &key, &vel) /*|| te.Message.GetNoteOff(&ch, &key, &vel)*/ {
			notesOn = append(notesOn, key)
		}
	})

	return notesOn
}

// Coal v1.0
func lcs(song1 []uint8, song2 []uint8) []uint8 {
	m := len(song1)
	n := len(song2)

	// Create an (m+1) x (n+1) matrix to store the lengths of LCS
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	maxLength := 0
	endingIndex := 0

	// Build the dp matrix
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if song1[i-1] == song2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLength {
					maxLength = dp[i][j]
					endingIndex = i
				}
			} else {
				dp[i][j] = 0
			}
		}
	}

	// Extract the longest common subsequence
	if maxLength == 0 {
		return nil
	}

	longestSubsequence := song1[endingIndex-maxLength : endingIndex]
	return longestSubsequence
}
