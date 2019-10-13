var app = new Vue({
    el: '#app',
    data: {
        paginator: {
            total_page: 0,
            records: [],
            limit: 3,
            page: 1,
            prev_page: 0,
            next_page: 0,
        },

        editUser: null,
        newUser: {first_name: '', last_name: '', img: ''},
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
                    this.getUsers();
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
        createUser(newUser) {
            fetch("/users", {
                body: JSON.stringify(newUser),
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
            })
                .then(() => {
                    this.getUsers();
                    this.msg = newUser.first_name + " " + newUser.last_name + " has been created!"
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
        next() {
            if (this.paginator.next_page <= this.paginator.total_page) {
                this.paginator.prev_page = this.paginator.page++;
                this.paginator.next_page = this.paginator.page + 1;
                this.getUsers();
            } else {
                console.log("next page is" + (this.paginator.next_page + 1));
                console.log("total is:" + this.paginator.total_page);
            }

        },
        prev() {
            if (this.paginator.prev_page > 0) {
                this.paginator.next_page = this.paginator.page--;
                this.paginator.prev_page = this.paginator.page - 1;
                this.getUsers();
            } else {
                console.log("prev page is" + this.paginator.prev_page - 1);
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
        }
    },
    mounted() {
        this.getUsers();
    },
});