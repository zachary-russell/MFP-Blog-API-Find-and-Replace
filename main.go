  package main

  import (
	  "encoding/json"
	  "fmt"
	  "io/ioutil"
	  "log"
	  "net/http"
	  "regexp"
	  "strconv"
  )
  type BlogPost struct {
	  ID      int    `json:"id"`
	  Link string `json:"link"`
	  Title       struct {
		  Rendered string `json:"rendered"`
	  } `json:"title"`
	  Content struct {
		  Rendered  string `json:"rendered"`
	  } `json:"content"`
	  Excerpt struct {
		  Rendered  string `json:"rendered"`
	  } `json:"excerpt"`
  }
  func main() {
	allPosts := getPosts()
	matched := findKeywords(allPosts)
	for i := range matched {
		fmt.Println(
			matched[i].Link)
	  }

}

func findKeywords(posts []BlogPost) []BlogPost {
	var matches []BlogPost
	var re = regexp.MustCompile(`Under Armour Connected Fitness|underarmour|Under Armour| UA |Connected Fitness|UACF `)

	for i := range posts {
		if re.MatchString(posts[i].Title.Rendered) || re.MatchString(posts[i].Content.Rendered) {
			matches = append(matches, posts[i])
		}
	}
	return matches

}
func getPosts() []BlogPost {
	page := 1
	var posts []BlogPost
	for {
		resp, err := http.Get(fmt.Sprint("https://blog.myfitnesspal.com/wp-json/wp/v2/posts?per_page=100&page=", strconv.Itoa(page)))
		if err != nil {
			log.Fatalln(err)
		}
		// break on last page of posts
		if resp.StatusCode == 400 {
			break
		} else {
			var tmpPosts []BlogPost
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}
			json.Unmarshal(body, &tmpPosts)
			posts = append(posts, tmpPosts...)
		}
		page++
	}
	return posts
}