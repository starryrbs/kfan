
export function StoreUser(username, id) {
    localStorage.setItem("user", JSON.stringify({"name": username, "id": id}))
}

export function GetUser() {
    const item = localStorage.getItem("user")
    if (item) {
        return JSON.parse(item)
    }
}