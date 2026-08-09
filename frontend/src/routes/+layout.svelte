<script lang="ts">
  import "./layout.css";

  let { children } = $props();
</script>

<svelte:head>
  <title>autopost</title>

  <meta
    name="description"
    content="Manage your automated social media posts."
  />
</svelte:head>
{@render children()}
