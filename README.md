## Welcome to GitHub Pages

You can use the [editor on GitHub](https://github.com/sanketpayghan/nolan/edit/master/README.md) to maintain and preview the content for your website in Markdown files.

Whenever you commit to this repository, GitHub Pages will run [Jekyll](https://jekyllrb.com/) to rebuild the pages in your site, from the content in your Markdown files.

### Markdown

Markdown is a lightweight and easy-to-use syntax for styling your writing. It includes conventions for

```markdown
Syntax highlighted code block

# Header 1
## Header 2
### Header 3

- Bulleted
- List

1. Numbered
2. List

**Bold** and _Italic_ and `Code` text

[Link](url) and ![Image](src)
```

For more details see [GitHub Flavored Markdown](https://guides.github.com/features/mastering-markdown/).

### Jekyll Themes

Your Pages site will use the layout and styles from the Jekyll theme you have selected in your [repository settings](https://github.com/sanketpayghan/nolan/settings). The name of this theme is saved in the Jekyll `_config.yml` configuration file.

### Support or Contact

Nolan helps you make efficient concurrent API calls or RPC from python.

Following functions can be used for above mentioned functionality:

1. `call_parallel_api(url, data_str, concurrent_on)`:
    - url - API Url (method will be post, will add more methods).
    - data_str - request parameters that needs to pass to request in string format(json.dumps can also be used)
    - concurrent_on - comma seprated primary keys or random number. API will be called that many times concurrently.
    
    Respose will be in string format concatenated for all the API requests. (Currently working on it to merge response     instead of simply concating responses of diffrent API calls.

2. `call_parallel_rpc(url, data_str, concurrent_on)`:
    - URL - Url for RPC
    - data_str - data to be passed in string format.
    - concurrent_on - comma seprated primary keys or random number. API will be called that many times concurrently.
    
    Respose will be in string format concatenated for all the API requests. (Currently working on it to merge response     instead of simply concating responses of diffrent API calls.

3. `parallel_rpc_with_data(url, data_str_list)`:
    - URL - Url for RPC
    - data_str_list - will contain data for each concurrent call. RPC will be called concurrently for all items in list with given values as parameters.
    
    Respose will be in string format concatenated for all the API requests. (Currently working on it to merge response     instead of simply concating responses of diffrent API calls.


Having trouble with Pages? Check out our [documentation](https://help.github.com/categories/github-pages-basics/) or [contact support](https://github.com/contact) and we’ll help you sort it out.
