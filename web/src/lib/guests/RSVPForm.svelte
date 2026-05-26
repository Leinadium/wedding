<script lang="ts">
  import GuestRow from "./GuestRow.svelte";

  // State
  let isAttending = true; // Default: I will go
  let isSubmitting = false;
  let feedbackMessage = "";

  // Array holding the list of guests
  let guests = [{ name: "", phone: "", isChild: false }];

  function addGuest() {
    guests = [...guests, { name: "", phone: "", isChild: false }];
  }

  function removeGuest(index: number) {
    guests = guests.filter((_, i) => i !== index);
  }

  async function handleSubmit(event) {
    event.preventDefault();
    isSubmitting = true;
    feedbackMessage = "";

    // Filter out empty rows just in case, though the 'required' HTML attribute helps prevent this
    const payload = guests.filter((g) => g.name.trim() !== "");

    if (payload.length === 0) {
      feedbackMessage = "Please provide at least one valid guest name.";
      isSubmitting = false;
      return;
    }

    // Determine route based on attendance
    const endpoint = isAttending ? "/v1/confirmation" : "/v1/rejection";

    try {
      const response = await fetch(endpoint, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(payload),
      });

      if (response.ok) {
        feedbackMessage = isAttending
          ? "Fantastic! We can't wait to see you."
          : "We're sorry you can't make it. Thank you for letting us know!";

        // Optional: Reset form after successful submission
        guests = [{ name: "", phone: "", isChild: false }];
        isAttending = true;
      } else {
        feedbackMessage =
          "Something went wrong saving your RSVP. Please try again.";
      }
    } catch (error) {
      console.error("RSVP Submission Error:", error);
      feedbackMessage =
        "Network error. Please check your connection and try again.";
    } finally {
      isSubmitting = false;
    }
  }
</script>

<div class="rsvp-container">
  <h2>RSVP</h2>

  {#if feedbackMessage}
    <div class="feedback">
      {feedbackMessage}
    </div>
  {/if}

  <form on:submit={handleSubmit}>
    <fieldset class="attendance-selector">
      <legend>Will you be joining us?</legend>
      <label class="radio-label">
        <input type="radio" bind:group={isAttending} value={true} />
        I will go!
      </label>
      <label class="radio-label">
        <input type="radio" bind:group={isAttending} value={false} />
        I won't go
      </label>
    </fieldset>

    <div class="guests-list">
      {#each guests as guest, i}
        <GuestRow
          bind:guest={guests[i]}
          onRemove={guests.length > 1 ? () => removeGuest(i) : null}
        />
      {/each}
    </div>

    <button type="button" class="add-guest-btn" on:click={addGuest}>
      + More guests include in the invite? Add them here
    </button>

    <button type="submit" class="submit-btn" disabled={isSubmitting}>
      {isSubmitting ? "Sending..." : "Send RSVP"}
    </button>
  </form>
</div>

<style>
  .rsvp-container {
    max-width: 600px;
    margin: 0 auto;
    padding: 2rem;
    background: #fff;
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  }
  .attendance-selector {
    border: none;
    padding: 0;
    margin-bottom: 2rem;
    display: flex;
    gap: 2rem;
  }
  .radio-label {
    font-weight: bold;
    cursor: pointer;
  }
  .guests-list {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    margin-bottom: 1.5rem;
  }
  .add-guest-btn {
    background: none;
    color: #0066cc;
    border: 1px dashed #0066cc;
    padding: 0.75rem;
    width: 100%;
    border-radius: 4px;
    cursor: pointer;
    margin-bottom: 2rem;
    font-weight: 500;
    transition: background 0.2s;
  }
  .add-guest-btn:hover {
    background: #f0f7ff;
  }
  .submit-btn {
    width: 100%;
    padding: 1rem;
    background-color: #2c3e50;
    color: white;
    border: none;
    border-radius: 4px;
    font-size: 1.1rem;
    cursor: pointer;
  }
  .submit-btn:disabled {
    background-color: #95a5a6;
    cursor: not-allowed;
  }
  .feedback {
    padding: 1rem;
    margin-bottom: 1rem;
    background-color: #e8f8f5;
    border-left: 4px solid #1abc9c;
  }
</style>
