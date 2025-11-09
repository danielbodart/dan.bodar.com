---
title: "GitHub Get Public Keys for a User"
date: 2018-02-08T10:58:49Z
slug: "github-get-public-keys-for-a-user"
categories:
  - Uncategorized
---

<pre><code>https://api.github.com/users/danielbodart/keys</code></pre>

Will return 
<pre><code>[
  {
    "id": 11647754,
    "key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDhoeoQelGjp0Oz3m64SV8AszPTyxlK1jJ8lYlW28WaSK5hu79BemIibekO06P2aA47vhI8hc4R90aqE0Ja1qG6A+jkezn2jeSzZZq7/dvoYFC0i96ljlJTKoOY5E0P4/259giTL1wtMbJrYfFutmVMBidHiOzA1dnLTrZthOuYGlanZL7oazfrz4Ht+ZdvoMoIFqU8yRnW/E903dn2DjH6Wk2dhsdP1QssbOtu3BNguhpcuVksrKw7hqrxCpiJW7Eo7z8yg6qP+KbzGF6G4UmA1/VDn5RFwNC7bvSShirYasEyAOrMHcUiCyvvvtT0a+SlsVfYbyH8wm5Ci38D9HK1"
  },
  {
    "id": 26992755,
    "key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDJDZbFgT04IDx5SECPFnHlJfwq/aHkhGr2sgSl0gqo3fOy48du/DPy4VH5cAk0qiLbT8dfzgMmkQm8EourUidZHORlbfbzzptMhWnbwEQbfXBuVyVa1Us5MV6Y0IV0TDOjNIALi6z3QTqJPtoOOfuRnDoumGCnl7bqxaKNrreoVm2vc7ep6yKuBR9HqDBJ6cN6KZ08IYFWjora1CscwsKC2EP6LWcnLsLVh2x4ZgJI0mDtnQ8hHedWmvQGhcyctKg1WdtcrrsVQESwlwDCiEVK9MypThYdZfVhh8d0JTApw37+QVkDW/oQeSMkKEyfss3wuOCCoGiJmurwM18Q/6V7"
  }
]
</code></pre>

Should return them in the same order as https://github.com/settings/keys

