document.getElementById("form").addEventListener("submit", async function(event) {
    event.preventDefault();
    auth = document.getElementById("auth").value
    body = document.getElementById("body").value
    method = document.getElementById("method").value
    path = document.getElementById("path").value

    let response = await fetch("https://" + window.location.hostname + path, {
        method: method,
        headers: {
            "Authorization": `Bearer ${auth}`
        },
        body: method == "GET" || "HEAD" ? undefined : body 
    })

    document.getElementById("out").innerHTML = await response.text()
}, true)