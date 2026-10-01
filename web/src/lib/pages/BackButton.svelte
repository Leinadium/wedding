<script lang="ts">
  import Text from "../text/Text.svelte";

  let {
    textKey,
    color,
    action,
  }: {
    textKey: string;
    color: string;
    action: () => void;
  } = $props();
</script>

<div style="--theme-color: {color}">
  <a
    class="formal"
    href="/#"
    onclick={(e) => {
      e.preventDefault();
      action();
    }}
  >
    <Text key={textKey} />
  </a>
</div>

<style>
  div {
    text-align: center;
    margin-bottom: 2rem;
    display: inline-block;
    margin-top: 2rem;
  }

  a {
    text-decoration: none;
    color: var(--theme-color);
    font-weight: 300;
    border: 1px solid var(--theme-color);
    border-radius: 10px;
    padding: 1rem;
  }
</style>
