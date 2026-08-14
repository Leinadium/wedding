<script lang="ts">
  import { getText } from "./text";

  let { key }: { key: string } = $props();
  let content: string = $derived(getText(key));
</script>

{@html content}
