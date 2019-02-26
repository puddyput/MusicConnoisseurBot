package musicData

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"strings"
)

type Database struct {
	DB *bolt.DB
}

func Init(dbFile string) *Database {
	var err error
	mdb := Database{}
	mdb.DB, err = bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Verify all Top-Level-Buckets exist
	err = mdb.DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Tracks"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &mdb
}
func (mdb Database) Close() {
	mdb.DB.Close()
}

func (mdb Database) PutTrack(key string, track Track) {
	encoded, err := track.Serialize()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = mdb.DB.Update(func(tx *bolt.Tx) error {
		key := []byte(track.URL)
		tx.Bucket([]byte("Tracks")).Put(key, encoded)
		return nil
	})
	if err != nil {
		log.Fatal(err)
		return
	}
}

func (mdb Database) Find(search string) []Track {
	var tracks []Track
	err := mdb.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tracks"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var t Track
			err := json.Unmarshal(v, &t)
			if err != nil {
				log.Print(err)
			}
			// check hashtags
			if Contains(t.Hashtags, search) {
				tracks = append(tracks, t)
				continue
			}
			// check title & description
			if strings.Contains(t.Title, search) ||
				strings.Contains(t.Description, search) {
				tracks = append(tracks, t)
				continue
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return tracks
}

func (mdb Database) FindByMessageId(id int) Track {
	var match Track
	err := mdb.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tracks"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var t Track
			err := json.Unmarshal(v, &t)
			if err != nil {
				log.Print(err)
			}
			if t.MessageId == id {
				match = t
				return nil
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return match
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
