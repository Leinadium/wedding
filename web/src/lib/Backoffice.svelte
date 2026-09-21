<script lang="ts">
  import { api, type Attendee } from "./api";

  let isAuthenticated: boolean = $state(false);
  let value: string | null = $state(null);
  let error: boolean = $state(false);
  let content = $state<Attendee[]>([]);

  async function handleLogin(e: Event) {
    e.preventDefault();
    if (value) {
      let res = await api.getAttendees(value);
      if (res) {
        error = false;
        content = res.attendees;
        isAuthenticated = true;
      } else {
        error = true;
        isAuthenticated = false;
      }
    }
  }

  async function copyToClipboard(text: string) {
    await navigator.clipboard.writeText(text);
  }
</script>

<svelte:head>
  <link
    rel="stylesheet"
    href="https://assets.ubuntu.com/v1/vanilla_framework_version_4.59.0.min.css"
  />
</svelte:head>

<div class="l-main">
  {#if !isAuthenticated}
    <section class="p-strip">
      <div class="row">
        <div class="col-4 prefix-4">
          <div class="p-card">
            <h3>Authentication</h3>
            <hr />
            <form class="p-form" onsubmit={handleLogin}>
              <fieldset>
                <label for="pass">Password</label>
                <input
                  id="pass"
                  type="password"
                  placeholder="Enter password..."
                  bind:value
                  required
                />
                {#if error}
                  <p
                    class="p-form-validation__message"
                    id="exampleInputErrorMessage"
                  >
                    Could not login.
                  </p>
                {/if}
                <button
                  type="submit"
                  class="p-button--positive u-no-margin--bottom"
                >
                  Unlock
                </button>
              </fieldset>
            </form>
          </div>
        </div>
      </div>
    </section>
  {:else}
    <section class="p-strip is-shallow">
      <div class="row">
        <div class="col-12">
          <h4 class="p-heading--4">Attendees</h4>

          <table aria-label="Attendees table">
            <thead>
              <tr>
                <th aria-sort="none">Name</th>
                <th aria-sort="none">Status</th>
                <th aria-sort="none" class="u-align--left">Updated At</th>
                <th>ID</th>
                <th>Invite ID</th>
              </tr>
            </thead>
            <tbody>
              {#each content as at}
                <tr>
                  <td>{at.name}</td>
                  <td>
                    {#if at.confirmed}
                      <span
                        class="p-chip--positive is-readonly is-inline is-dense"
                      >
                        Confirmed
                      </span>
                    {:else if at.confirmed === false}
                      <span
                        class="p-chip--negative is-readonly is-inline is-dense"
                      >
                        Won't go
                      </span>
                    {:else}
                      <span
                        class="p-chip--caution is-readonly is-inline is-dense"
                      >
                        To be confirmed
                      </span>
                    {/if}
                  </td>

                  <td class="u-align--left">{at.updatedAt}</td>

                  <!-- Truncated UUID with Copy Button -->
                  <td class="uuid-cell">
                    <div class="uuid-container">
                      <!-- u-truncate automatically handles the "..." overflow styling -->
                      <code class="u-truncate" title={at.id}>
                        {at.id}
                      </code>
                      <button
                        class="copy-btn"
                        aria-label="Copy ID to clipboard"
                        onclick={() => copyToClipboard(at.id)}
                      >
                        <i class="p-icon--copy"></i>
                      </button>
                    </div>
                  </td>

                  <!-- Truncated Invite ID with Copy Button -->
                  <td class="uuid-cell">
                    <div class="uuid-container">
                      <code class="u-truncate" title={at.inviteId}>
                        {at.inviteId}
                      </code>
                      <button
                        class="copy-btn"
                        aria-label="Copy Invite ID to clipboard"
                        onclick={() => copyToClipboard(at.inviteId)}
                      >
                        <i class="p-icon--copy"></i>
                      </button>
                    </div>
                  </td>
                </tr>
              {/each}
            </tbody>
          </table>
        </div>
      </div>
    </section>
  {/if}
</div>
