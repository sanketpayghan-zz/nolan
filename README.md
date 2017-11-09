
#**Nolan**


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
