<script>
  // The guest object passed from the parent
  export let guest = { name: "", phone: "", isChild: false };

  // Optional remove function (passed from parent if there are multiple rows)
  export let onRemove = null;
</script>

<div class="guest-row">
  <div class="input-group">
    <label for="name">Name</label>
    <input
      type="text"
      id="name"
      bind:value={guest.name}
      placeholder="John Doe"
      required
    />
  </div>

  <div class="input-group">
    <label for="phone">Phone</label>
    <input
      type="tel"
      id="phone"
      bind:value={guest.phone}
      placeholder="+1 234 567 8900"
      required
    />
  </div>

  <div class="checkbox-group">
    <label>
      <input type="checkbox" bind:checked={guest.isChild} />
      Is Under 12
    </label>
  </div>

  {#if onRemove}
    <button
      type="button"
      class="remove-btn"
      on:click={onRemove}
      aria-label="Remove guest"
    >
      ✕
    </button>
  {/if}
</div>

<style>
  .guest-row {
    display: flex;
    flex-wrap: wrap;
    gap: 1rem;
    align-items: flex-end;
    padding: 0.5rem 0;
    border-bottom: 1px solid #eee;
  }
  .input-group {
    display: flex;
    flex-direction: column;
    flex: 1;
    min-width: 150px;
  }
  .checkbox-group {
    display: flex;
    align-items: center;
    height: 40px; /* Aligns with inputs */
  }
  input[type="text"],
  input[type="tel"] {
    padding: 0.5rem;
    border: 1px solid #ccc;
    border-radius: 4px;
  }
  .remove-btn {
    background: transparent;
    color: #ff4444;
    border: none;
    cursor: pointer;
    font-size: 1.2rem;
    padding: 0.5rem;
  }
</style>
