<script lang="ts">
    import clsx from 'clsx';

    export let appearance: 'primary' | 'secondary' = 'primary';
    export let size: 'xs' | 'sm' | 'md' | 'lg' | 'xl' = 'md';

    const appearances = {
        primary: 'bg-indigo-500 text-white shadow-sm hover:bg-indigo-400 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-500',
        secondary: ''
    }

    const sizes = {
        xs: 'px-2 py-1 text-xs font-semibold',
        sm: 'px-2 py-1 text-sm font-semibold',
        md: 'px-2.5 py-1.5 text-sm font-semibold',
        lg: 'px-3 py-2 text-sm font-semibold',
        xl: 'px-3.5 py-2.5 text-sm font-semibold',
    }
</script>


<button
    {...$$props}
    class={clsx(
        'rounded',
        appearances[appearance],
        sizes[size], 
        $$props.class,
    )}
>
    <slot />
</button>
<!-- <button type="button" class="rounded bg-indigo-500 px-2 py-1 text-xs font-semibold text-white shadow-sm hover:bg-indigo-400 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-500">Button text</button>
<button type="button" class="rounded bg-indigo-500 px-2 py-1 text-sm font-semibold text-white shadow-sm hover:bg-indigo-400 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-500">Button text</button>
<button type="button" class="rounded-md bg-indigo-500 px-2.5 py-1.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-400 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-500">Button text</button>
<button type="button" class="rounded-md bg-indigo-500 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-400 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-500">Button text</button>
<button type="button" class="rounded-md bg-indigo-500 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-400 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-500">Button text</button> -->