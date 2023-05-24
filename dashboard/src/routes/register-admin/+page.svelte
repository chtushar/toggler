<script lang="ts">
    import Button from '$lib/ui/Button.svelte';
    import axios from '../../utils/axios'

    let password = '';
    let confirmPassword = '';

    const handleAddAdmin = async (e: Event & { currentTarget: EventTarget & HTMLFormElement; }) => {
        e.preventDefault();

        const formData = new FormData(e.target as HTMLFormElement);
        await axios.post('/api/add_admin', {
            email: formData.get('email'),
            name: formData.get('name'),
            password: formData.get('password')
    })
}
</script>

<div class="flex w-full h-full flex-col items-center justify-center mx-auto max-w-xs">
    <form class="w-full flex flex-col gap-8" on:submit={handleAddAdmin}>
        <div class="flex flex-col space-y-4">
            <div class="">
                <label for="email-input">Email</label>
                <input id="email-input" class="mt-2" type="text" name="email" required aria-required="true" />
            </div>
            <div>
                <label for="name-input">Name</label>
                <input type="text" class="mt-2" id="name-input" name="name" required aria-required="true" />
            </div>
            <div>
                <label for="password-input">Password</label>
                <input type="password" bind:value={password} id="password-input" class="mt-2" name="password" required aria-required="true" />
            </div>
            <div>
                <label for="confirm-password-input">Confirm Password</label>
                <input type="password" bind:value={confirmPassword} id="confirm-password-input" class="mt-2" name="password" required aria-required="true" />
            </div>
        </div>
        <Button type="submit" class="w-full" size="xl" disabled={password !== confirmPassword} >Add admin</Button>
    </form>
</div>
