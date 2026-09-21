<script lang="ts">
  import type { Attendee } from "../../api";
  import Row from "./Row.svelte";

  let {
    content,
  }: {
    content: Attendee[];
  } = $props();

  // Explicitly define only the keys you want to be sortable
  type SortKey = "name" | "confirmed" | "updatedAt";
  type SortDirection = "none" | "ascending" | "descending";

  let sortKey = $state<SortKey | null>(null);
  let sortDirection = $state<SortDirection>("none");

  function sortBy(key: SortKey) {
    if (sortKey === key) {
      if (sortDirection === "ascending") {
        sortDirection = "descending";
      } else if (sortDirection === "descending") {
        sortDirection = "none";
        sortKey = null;
      } else {
        sortDirection = "ascending";
      }
    } else {
      sortKey = key;
      sortDirection = "ascending";
    }
  }

  let sortedContent = $derived.by(() => {
    if (sortDirection === "none" || sortKey === null) {
      return content;
    }

    const key = sortKey;

    return [...content].sort((a, b) => {
      // Cast to 'any' here since we removed keyof Attendee.
      // This stops TypeScript from complaining about the index type.
      const valA = String((a as any)[key] || "");
      const valB = String((b as any)[key] || "");

      return sortDirection === "ascending"
        ? valA.localeCompare(valB)
        : valB.localeCompare(valA);
    });
  });
</script>

<table aria-label="Attendees table">
  <thead>
    <tr>
      <th aria-sort={sortKey === "name" ? sortDirection : "none"}>
        <button class="p-table__sort-button" onclick={() => sortBy("name")}>
          Name
        </button>
      </th>
      <th aria-sort={sortKey === "confirmed" ? sortDirection : "none"}>
        <button
          class="p-table__sort-button"
          onclick={() => sortBy("confirmed")}
        >
          Confirmed
        </button>
      </th>
      <th aria-sort={sortKey === "updatedAt" ? sortDirection : "none"}>
        <button
          class="p-table__sort-button"
          onclick={() => sortBy("updatedAt")}
        >
          Updated At
        </button>
      </th>
      <th>Invite ID</th>
    </tr>
  </thead>
  <tbody>
    {#each sortedContent as at}
      <Row {at} />
    {/each}
  </tbody>
</table>
