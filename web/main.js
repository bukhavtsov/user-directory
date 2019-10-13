var app = new Vue({
    el: '#app',
    data: {
        paginator: {
            total_page: 0,
            records: [],
            limit: 3,
            page: 1,
        },

        editUser: null,
        user: {first_name: '', last_name: '', img: ''},
        users: [],
        img: null,
        msg: ""
    },
    methods: {
        deleteUser(id, i) {
            fetch("/users/" + id, {
                method: "DELETE",
                headers: {
                    "Content-Type": "application/json",
                }
            })
                .then(() => {
                    this.users.splice(i, 1);
                })
        },
        updateUser(user) {
            fetch("/users/" + user.id, {
                body: JSON.stringify(user),
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                },
            })
                .then(() => {
                    this.editUser = null;
                })
        },
        createUser(user) {
            fetch("/users", {
                body: JSON.stringify(user),
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
            })
                .then(() => {
                    this.getUsers();
                    this.msg = user.first_name + " " + user.last_name + " has been created!"
                });
        },
        upload(event, user) {
            let fd = new FormData();
            let file = event.target.files[0];

            fd.append('name', 'user_icon');
            fd.append('file', file);

            let config = {
                header: {
                    'Content-Type': 'multipart/form-data'
                }
            };

            //FIXME: rewrite without axios or write axios for all functions
            axios.put('/users/uploadicon/' + user.id, fd, config)
                .then(() => {
                    this.editUser = null;
                    this.getUsers();
                });
        },
        setPage(page) {
            console.log("page is:" + page)
            if (page > 0 && page <= this.paginator.total_page) {
                this.paginator.page = page;
                this.getUsers();
            }

        },
        getUsers() {
            fetch("/users/pagination/" + this.paginator.page + "/" + this.paginator.limit)
                .then(response => response.json())
                .then((data) => {
                    this.users = data.records;
                    this.paginator.total_page = data.total_page;
                    this.paginator.page = data.page;
                })
        },
        find(user) {
            fetch("/users/find/" + user.first_name + "/" + user.last_name)
                .then(response => response.json())
                .then((data) => {
                    this.users = data;
                })
        },
        clickCallback (page)  {
            console.log(page)
        },

    },
    mounted() {
        this.getUsers();
    },
});