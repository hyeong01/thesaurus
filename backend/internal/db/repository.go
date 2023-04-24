package db

func GetSynonymsAndSimilarities(word string) ([]string, []int, error) {
	query := "SELECT * FROM synonyms"
	rows, err := conn.Query(query, word)
	print(rows)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var synonyms []string
	var similarities []int

	for rows.Next() {
		var synonym string
		var similarity int
		err = rows.Scan(&synonym, &similarity)
		if err != nil {
			return nil, nil, err
		}
		synonyms = append(synonyms, synonym)
		similarities = append(similarities, similarity)
	}

	err = rows.Err()
	if err != nil {
		return nil, nil, err
	}

	return synonyms, similarities, nil
}
