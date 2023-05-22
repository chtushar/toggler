<script lang="ts">
    import axios from '../../utils/axios'

    const handleAddAdmin = async (e: Event & { currentTarget: EventTarget & HTMLFormElement; }) => {
        const formData = new FormData(e.target as HTMLFormElement);
        await axios.post('/api/add_admin', {
            email: formData.get('email'),
            name: formData.get('name'),
            password: formData.get('password')
    })
}
</script>

<div class="flex w-full h-full flex-col items-center justify-center mx-auto max-w-xs">
    <form class="w-full flex flex-col gap-6" on:submit={handleAddAdmin}>
        <div class="flex flex-col gap-3">
            <input type="text" placeholder="Email" name="email" required aria-required="true" />
            <input type="text" placeholder="Name" name="name" required aria-required="true" />
            <input type="password" placeholder="Password" name="password" required aria-required="true" />
        </div>
        <button type="submit">Create</button>
    </form>
</div>
