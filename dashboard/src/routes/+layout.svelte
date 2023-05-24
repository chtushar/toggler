<script lang="ts">
	import '../app.scss';
	import { browser } from '$app/environment';
    import Layout from '$lib/common/Layout.svelte';
	import { QueryClient, QueryClientProvider, createQuery } from '@tanstack/svelte-query';

	const queryClient = new QueryClient({
		defaultOptions: {
			queries: {
				enabled: browser
			}
		}
	});
</script>

<QueryClientProvider client={queryClient}>
    <Layout>
        <slot />
    </Layout>
</QueryClientProvider>
