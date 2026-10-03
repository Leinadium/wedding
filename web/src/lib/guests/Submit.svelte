<script lang="ts">
  import { fade } from "svelte/transition";
  import { getText } from "../text/text";

  let {
    onClick,
    success,
    loading,
    isDisabled,
  }: {
    onClick: () => void;
    success: boolean;
    loading: boolean;
    isDisabled: boolean;
  } = $props();

  let disabled: boolean = $derived(loading || success || isDisabled);

  let text: string = $derived(
    getText(
      success ? "invite-saved" : loading ? "invite-saving" : "invite-save",
    ),
  );

  function handle() {
    if (!success) onClick();
  }
</script>

{#if !isDisabled || success}
  <input
    in:fade
    class={["save", "formal", { success }, { loading }, { disabled }]}
    type="submit"
    value={text}
    onclick={handle}
    {disabled}
  />
{/if}

<!-- <span class="warning formal">{getText("invite-warning")}</span> -->

<style>
  .save {
    padding: 0.5rem 1rem;
    background-color: transparent;
    color: #dedacd;
    border: none;
    border-radius: 6px;
    border: 3px solid #dedacd;
    font-size: 1.2rem;
    cursor: pointer;
    transition: background-color 0.2s;
  }

  .disabled {
    background-color: #aaa599;
    opacity: 50%;
  }

  .success {
    background-color: #777155;
  }

  .loading {
    background-color: #aaa599;
  }
  .warning {
    color: #dedacd;
    font-size: 1rem;
  }
</style>
