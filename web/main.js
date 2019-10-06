const app = new Vue({
    el: "#app",
    data: {
        editUser: null,
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
        }
    },
    mounted() {
        fetch("http://localhost:8080/users")
            .then(response => response.json())
            .then((data) => {
                this.users = data;
                console.log(data)
            })
    },
    template: `
    <div>
      <li v-for="user, i in users">
        <div v-if="editUser === user.id">
          <input v-on:keyup.13="updateUser(user)" v-model="user.name" />
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