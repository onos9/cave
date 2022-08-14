<script lang="ts">
	const form = (e: any) => {
		e.preventDefault();
		const formdata = new FormData(e.target);
		const data = Object.fromEntries(formdata);
		console.log(formdata);
	};
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
