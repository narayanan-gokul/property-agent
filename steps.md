1. Perform property search on [domain](https://domain.com.au).
2. Perform the following action on each result:
    1. Check available date.
        - Only select properties available on or after June 31st.
    2. Add to shortlist.
        - Extract property location.
        - Extract listing link.
        - Extrack availability date.

- Make request on the URL.
- Find the lines with the string `css-1y2bib4`.
- Find `href="*"`. `*` is the URL for the listing.
- Navigate to the URL.
- Find `Date Available: `. Everything after the result enclosed within
`<strong>*</strong>` is the availability date. Length 24
- `<h1 class="css-hkh81z">*</h1>` is the element with the address. Length 23
- `<div data-testid="listing-details__summary-title" class="css-twgrok"><span>*</span></div>` is the element with the price. Length 75
- Find `googleapis`. The numbers from `center=` until `&` is the location, separated by commas.
---

# Rough area:

---

URL:
/rent/?suburb=darlington-nsw-2008,redfern-nsw-2016,ultimo-nsw-2007,glebe-nsw-2037,surry-hills-nsw-2010,newtown-nsw-2042,pyrmont-nsw-2009,chippendale-nsw-2008,annandale-nsw-2038,eveleigh-nsw-2015,stanmore-nsw-2048&bedrooms=2-any&bathrooms=2-any&price=0-1000&availableto=2025-07-14&excludedeposittaken=1&page=num
`"Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8"`
`"User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 Chrome/113.0.0.0 Safari/537.36"`
`"Cookies", "DEVICE_SESSIONID=47dc4841-b4ab-4c58-b3da-95fd1ed7537d; domain-mixpanel-id_ab0bde70050c3eabaaf8824402fa01e0=53458498; searchSOI=nsw; ppid=2bbde9d8f551ea5f9c2351e6ee9518b00590ca618c6d07e154bd6fac4a1bc6ce; ak_bmsc=6DBAFAD32509300BE19C847B79E578C1~000000000000000000000000000000~YAAQpZUvF4jXdC+XAQAA81q4ORxYRSVo3looQjwC8wGVgXkD++RLkz8etlsHWOMmSL8SvJj7DJZh5qAywx0TRPOUoRHK4WpNqyaNk4nbTVMz15KiQOcFXZU4J5gw0fyDNBGHhIm5QtDb3sn6cQjNibSYaRrcXpcee3NB5iwf4eXX9PapZ0J5D5BE6BiW+q3jiuDqyAnIkvbkjgbKN1ITxcIBehnksebQ/FXGSViraoPqRfa2o4zXGZf0FtFQvrhJwf55zKW3KDsy5v2aRs8+VCdG1/ahY0h45fGmNLn9jJTyCi8Vvf9kpxHijrAZWinsVgGpOicwIneIQP9NHmsQQbl1S0B2unk7ZkOQkX7tZmuZmUw2m7hLQ9VYZQoNHBFjOoQAGkJKj1TD/F9IBqyADNy8PmMjPfE8CvKxAnDlVnXqzMxZXUrXlT7S/dyeCsvyqt0J5REfEE0Fn6Gr1lNAVQ==; bm_sz=B066C0541C82443531FEF8E11DAD2DBF~YAAQpZUvF7cIeC+XAQAAPpjdORzxnM0/qhaLWdv4shqoy52X8zTv/0ZAvpTSOvS7vLNXjs1+qVll05DmJPb7zPcFyGG3Xf7SKh/E1HssOad+hog/IEzAl0mLQaD97rMJPMPiFLAff3ePrsU3kCalzWZqo+OWrZ1yJ1nEYNjnxeTJNcvewYS/ieq+uBGYQUcX1RX5qauXt0/jFm9UBCU99o9nc0EdcpVzU76Xx6g8flyGKyfuJAc1i434AztKBl5WZ3y8SU2/GEsVKWoF4mNeTGoo0Lm5YVrGB66SdL/BMm2cP4qyNuQe/fugYc5ce4HAeknTvXCl5gc7928Pzx96CL04080dd88cfZ6xHRegw97LJKf0IaJLlCypgEqtg2AcQ6w9l8MUrpCspjEj+AmXQ6iprP8CfPBmrRUr2QHlce5+juwhU7jDoNolLwOEIG8Mleg82E749pjHC2KpHezXyCkTC1Ql7ujBOWZF8F0RhHW/ORDH9rvs184g9imHDJVcj7MC594LT4CgbimgSiuTWuN6RVRIEg==~3424568~3749185; _abck=3F2E7042DAEBB7DCF4C6607BCCB29EC8~0~YAAQpZUvF1KNeS+XAQAAGkbvOQ4Z4bEcH+I/TmOyXd5TPjE5O8m4Zki52rRBckfXHyWQEKXetnwZ5q6/pDGM+MwYINFZ8yYdtCkFa68Ezm9+xFlpVhn6OcO9ZyIJ3EUroR5BpRX5Bb3Jg9DOyA9eQfqfKqvEVYmfUHYOdWB6zk30zlEbzMX6G+Yhd71A+8yn26sTZpnS3cVvxp44aW42qQAh8B0DA97gjlOUVFOI6cOswRtjc1ThLOosxP//2RbuxqFOIw9NZPdAQucnpMWjv2hz/xg3pXw2z+I+E0bTnAfNac8Cm5T16xw9gnacbQ8N/Zdn8oXuMcLqWbSZnjMd63zRGTuYTM+GkcZP2p1FA6Yv3Rusmn0v6EdbrqF7BMGSrcSEHZFzklik8pcdUrNWrLS7bkCLfpyV9KqA3cID3BSUvUJJqih0WxO84Tk16IC41qWrylceKy+PxWGPDjIYaYDtA1C4TgSCcPwKPPOopQKXbYwiwdqqOgfdEVXS9ZVCVt59W9Ib+pUTH2GMPm5Wic2lJ/QWrRxjThoLaHagdIjJe5m5bJVe+Ezx5v9F08vrA+YNgs0ZDkIU9xOBwd4pK4cUQQ+DvGkDd+P+xLf3INV0iPaq/qvLEzpotFfyV6DLwU6wqWH04imp/Qt5rWwjBRXBa7E/vg505pxB8/gSArK60V3szaofwGMhA/jKLlT0+0cEcwMDTcIyCMG6H3l1jJnrkF9ECgUHaR0y9TqIiN6srOBWEkD/h/0hhiLeGuUZVNsZhl/MRp6p1SfMKsEMIphatpvPwCJxiWie0/h6jNtOi6bCKppT+2VHYcjCAjg5/sEOBtfMVN5drplm78LmXRIp+be6XCznY1b4QLRN0h4nxOd5Q368wGsHeso/xl0+17bWnH54s9DSZ3EmF/M4AHomhMdQDvfZqvAKHUfMOZ8G60tMgsXueCHqI4zrZgM3~-1~-1~1749027271; bm_sv=6A62CD79E048AFFD6FA5AE2A7B6111E2~YAAQpZUvFxjgeS+XAQAAvpTzORzC3b9xsQYmvirmA0kL5s6g4Yr5WkpWM/8022iDONBE1djzD0ebM5w3BI2MIBdUnQ5oqL4roAVu0TqXETlBRskcnhv97ekpPzdTL6apdqDyho4MuwobmPdhWcLfAGrY9LcKLQ/pazjyKw7ElVt9DjXWGiuuGrNUOenPYLMoLiilE9b57yR0OrNRO8ZOje2+aXoWrXDuxe6KVX9tP4pXlN8H9BJ7FuAyjzZ+CZ4Zsz2szg==~1"`

