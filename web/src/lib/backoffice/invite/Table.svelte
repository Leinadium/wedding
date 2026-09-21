<script lang="ts">
  import type { Invite } from "../../api";
  import Row from "./Row.svelte";

  let {
    content,
  }: {
    content: Invite[];
  } = $props();
</script>

<table class="p-table--expanding" aria-label="Invites table">
  <thead>
    <tr>
      <th>ID</th>
      <th>Phone</th>
      <th>Attendees</th>
      <th>Note</th>
      <th aria-hidden="true"></th>
    </tr>
  </thead>
  <tbody>
    {#each content as invite}
      <Row {invite} />
    {/each}
  </tbody>
</table>
