<script lang="ts">
  import { api } from "../api";

  let name = $state("");
  let phone = $state("");
  let statusMessage = $state("");

  async function handleRSVP() {
    try {
      await api.registerGuest({ name, phone });
      statusMessage = "Success! See you there.";
    } catch (err: any) {
      statusMessage = `Error: ${err.message}`;
    }
  }
</script>

<div class="collapsible-content">
  <p>{statusMessage}</p>
  <form
    onsubmit={(e) => {
      e.preventDefault();
      handleRSVP();
    }}
  >
    <input bind:value={name} placeholder="Name" required />
    <input bind:value={phone} placeholder="Phone" required />
    <button type="submit">I'm coming!</button>
  </form>
</div>
