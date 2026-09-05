<script lang="ts">
  import { onMount } from "svelte";
  import type { Attendee } from "../api";
  import Text from "../text/Text.svelte";

  let {
    index,
    attendee,
    updateStatus,
  }: {
    index: number;
    attendee: Attendee;
    updateStatus: (index: number, status: boolean | null) => void;
  } = $props();

  let name: string = $derived(attendee.name);
  let selected: string = $state("pending");

  let status: boolean | null = $derived(attendee.confirmed);

  function handle() {
    console.log("handle");
    switch (selected) {
      case "confirmed":
        updateStatus(index, true);
        break;
      case "wontgo":
        updateStatus(index, false);
        break;
      default:
        updateStatus(index, null);
        break;
    }
  }

  onMount(() => {
    switch (status) {
      case true:
        selected = "confirmed";
        break;
      case false:
        selected = "wontgo";
        break;
      default:
        selected = "pending";
        break;
    }
  });
</script>

<div class="attendee">
  <span class="name formal {selected}">
    {attendee.name}
  </span>

  <label for="attendee-{name}"></label>
  <select
    bind:value={selected}
    onchange={handle}
    name="attendee"
    id="attendee-{name}"
    class="formal"
  >
    <option value="pending">
      <Text key="attendee-pending"></Text>
    </option>
    <option value="confirmed">
      <Text key="attendee-confirmed"></Text>
    </option>
    <option value="wontgo">
      <Text key="attendee-wontgo"></Text>
    </option>
  </select>

  <!-- <button class="status {statusClass}" onclick={toggleStatus}>
    {statusText}
  </button> -->
</div>

<style>
  .attendee {
    display: flex;
    flex-flow: row nowrap;
    justify-content: right;
    align-items: center;
    gap: 1rem;
  }

  .name {
    flex: 1;
    width: 100%;
    font-weight: 300;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    text-align: left;

    font-size: 1.3rem;
    color: #dedacd;
  }

  .wontgo {
    text-decoration: line-through;
  }

  select {
    border: none;
    border-bottom: 1px solid #fae7b6;

    color: #dedacd;
    font-size: 1rem;
    background: transparent;
  }
</style>
