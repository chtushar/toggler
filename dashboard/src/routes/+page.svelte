<script lang="ts">
	import axios from '../utils/axios'

	const handleAddAdmin = async (e: Event & { currentTarget: EventTarget & HTMLFormElement; }) => {
		const formData = new FormData(e.target as HTMLFormElement);
		await axios.post('/api/add_admin', {
			email: formData.get('email'),
			name: formData.get('name'),
			password: formData.get('password')
		})
	}

	const handleLogin = async (e: Event & { currentTarget: EventTarget & HTMLFormElement; }) => {
		const formData = new FormData(e.target as HTMLFormElement);
		await axios.post('/api/login', {
			email: formData.get('email'),
			password: formData.get('password')
		})
	}

	const handleGetAllUsers = async () => {
		const { data } = await axios.get('/api/get_users')
		console.log(data)
	}

	const handleLogout = async () => {
		await axios.post('/api/logout')
	}
</script>

<!-- Take email, name and password and create an admin -->
<form on:submit={handleAddAdmin}>
	<input type="text" placeholder="Email" name="email" />
	<input type="text" placeholder="Name" name="name" />
	<input type="password" placeholder="Password" name="password" />
	<button type="submit">Create</button>
</form>

<br />
<br />

<form on:submit={handleLogin}>
	<input type="text" placeholder="Email" name="email" />
	<input type="password" placeholder="Password" name="password" />
	<button type="submit">Login</button>
</form>

<br />
<br />

<button on:click={handleGetAllUsers}>Get all users</button>

<br />
<br />

<button on:click={handleLogout}>Logout</button>
