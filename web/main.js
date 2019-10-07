const app = new Vue({
    el: "#app",
    data: {
        editUser: null,
        newUser: {first_name: '', last_name: '', img: ''},
        users: [],
    },
    methods: {
        deleteUser(id, i) {
            fetch("http://localhost:8080/users/" + id, {
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
            fetch("http://localhost:8080/users/" + user.id, {
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
            fetch("http://localhost:8080/users", {
                body: JSON.stringify(newUser),
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
            })
                .then(() => {
                    fetch("http://localhost:8080/users")
                        .then(response => response.json())
                        .then((data) => {
                            this.users = data;
                        })
                })
        },
    },
    mounted() {
        fetch("http://localhost:8080/users")
            .then(response => response.json())
            .then((data) => {
                this.users = data;
            })
    },
    template: `
    <div>
    <input v-on:keyup.13="createUser(newUser)" v-model="newUser.first_name"/>
    <input v-on:keyup.13="createUser(newUser)" v-model="newUser.last_name"/>
    <input v-on:keyup.13="createUser(newUser)" v-model="newUser.img"/>
    <button v-on:click="createUser(newUser)">create</button>
      <li v-for="user, i in users">
        <div v-if="editUser === user.id">
          <input v-on:keyup.13="updateUser(user)" v-model="user.first_name"/>
          <input v-on:keyup.13="updateUser(user)" v-model="user.last_name" />
          <input v-on:keyup.13="updateUser(user)" v-model="user.img"  />
          <button v-on:click="updateUser(user)">save</button>
        </div>
        <div v-else>
          <button v-on:click="editUser = user.id">edit</button>
          <button v-on:click="deleteUser(user.id, i)">x</button>
          {{user.id}}
          {{user.first_name}}
          {{user.last_name}}
          {{user.img}}
        </div>
      </li>
    </div>
    `,
});