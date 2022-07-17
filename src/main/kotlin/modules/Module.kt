package modules

import dev.kord.core.Kord
import dev.kord.core.behavior.interaction.respondEphemeral
import dev.kord.core.entity.interaction.GuildChatInputCommandInteraction
import dev.kord.core.event.interaction.GuildChatInputCommandInteractionCreateEvent
import dev.kord.core.on
import java.time.LocalDateTime
import java.time.format.DateTimeFormatter

/**
 * Module
 * @param client The client instance.
 * @param name The name of the module.
 * @param commands The command names to listen for.
 */
open class Module
constructor(
        private val client: Kord,
        private val name: String,
        private val commands: Array<String>? = null,
) {
    init {
        moduleSetup()
    }

    /** Get module stats from file and log creation to console */
    private fun moduleSetup() {
        println("[Modules] Loading $name...")
        setupModule()
        if (commands !== null) {
            client.on<GuildChatInputCommandInteractionCreateEvent> {
                val command = interaction.command

                if (commands.contains(command.data.name.value)) {
                    println(
                            "[Commands] Executing /${command.data.name.value} | \uD83E\uDDCD ${interaction.user.username}#${interaction.user.discriminator} | \uD83D\uDDD3 ${
                            LocalDateTime.now().format(
                                DateTimeFormatter.ofPattern("dd/MM/yyyy HH:mm:ss")
                            )
                        }"
                    )
                    onCommand(interaction)
                }
            }
        }
        println("[Modules] Loaded $name!")
    }

    /** Called on startup */
    open fun setupModule() {}

    /** Called on Shutdown */
    open fun shutdown() {}

    /** Called when a chat input interaction is received and matches a command name. */
    open suspend fun onCommand(req: GuildChatInputCommandInteraction) {
        req.respondEphemeral { content = "This command isn't implemented!" }
    }
}
