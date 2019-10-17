package main

import "errors"

func queryMeme(memeID uint32) (data * memeInfo, err error){

	query, err := dbInfo.db.Query("SELECT id, link, post_time, stealer FROM memes WHERE id = ?", (memeID))

	if err == nil{

		var data = new(memeInfo);

		err := query.Scan(&data.ID, &data.link, &data.postTime, &data.creator);

		if err == nil{
			return data, nil;
		}
	}
	return nil, errors.New("Query Failed")
}

func querySize()(data *sizeResponse, err error){
	query, err := dbInfo.db.Query("SELECT COUNT(*) FROM memes;")

	if err == nil{

		var data = new (sizeResponse);

		err := query.Scan(&data.size)

		if err == nil{
			return data, nil
		}
	}

	return nil, errors.New("Query Failed")
}

func queryByTag(tags []string)(data * searchResponse, err error){

	query, err := dbInfo.db.Query("SELECT memes.id WHERE tags.tag = ? AND memes.id = tags.id", tags)

	if err == nil{

		var data = new(searchResponse)

		err := query.Scan(&data.ids)

		if err == nil{
			return data, nil
		}
	}
	return nil, errors.New("Query Failed")
}