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

	const handleCreateUser = async (e: Event & { currentTarget: EventTarget & HTMLFormElement; }) => {
		const formData = new FormData(e.target as HTMLFormElement);
		await axios.post('/api/create_user', {
			email: formData.get('email'),
			name: formData.get('name'),
			password: formData.get('password')
		})
	}

	const handleCreateProject = async (e: Event & { currentTarget: EventTarget & HTMLFormElement; }) => {
		const formData = new FormData(e.target as HTMLFormElement);
		await axios.post('/api/create_project')
	}	

	const handleGetAllUsers = async () => {
		const { data } = await axios.get('/api/get_users')
		console.log(data)
	}

	const handleLogout = async () => {
		await axios.post('/api/logout')
	}
</script>

<h1>
	Add admin
</h1>
<!-- Take email, name and password and create an admin -->
<form on:submit={handleAddAdmin}>
	<input type="text" placeholder="Email" name="email" />
	<input type="text" placeholder="Name" name="name" />
	<input type="password" placeholder="Password" name="password" />
	<button type="submit">Create</button>
</form>

<br />
<br />

<h1>
	Login
</h1>
<form on:submit={handleLogin}>
	<input type="text" placeholder="Email" name="email" />
	<input type="password" placeholder="Password" name="password" />
	<button type="submit">Login</button>
</form>

<br />
<br />

<h1>
	Create user
</h1>

<form on:submit={handleCreateUser}>
	<input type="text" placeholder="Email" name="email" />
	<input type="text" placeholder="Name" name="name" />
	<button type="submit">Create</button>
</form>

<br />
<br />


<h1>
	Create Project
</h1>

<form on:submit={handleCreateProject}>
	<input type="text" placeholder="Name" name="name" />
	<button type="submit">Create</button>
</form>

<br />
<br />

<button on:click={handleGetAllUsers}>Get all users</button>

<br />
<br />

<button on:click={handleLogout}>Logout</button>
