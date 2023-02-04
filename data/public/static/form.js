document.getElementById("form").addEventListener("submit", async function(event) {
    event.preventDefault();
    auth = document.getElementById("auth").value
    body = document.getElementById("body").value
    method = document.getElementById("method").value

    let response = await fetch("https://" + window.location.hostname, {
        method: method,
        headers: {
            "Authorization": `Bearer ${auth}`
        },
        body: body 
    })

    document.getElementById("out").innerHTML = await response.text()
}, true)