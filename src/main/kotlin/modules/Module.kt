package modules

import org.javacord.api.DiscordApi
import org.javacord.api.entity.message.MessageFlag
import org.javacord.api.interaction.SlashCommandInteraction
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
    private val client: DiscordApi,
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
            client.addSlashCommandCreateListener {
                val interaction = it.slashCommandInteraction

                if (commands.contains(interaction.commandName)) {
                    println(
                        "[Commands] Executing /${interaction.commandName} | \uD83E\uDDCD ${interaction.user.name}#${interaction.user.discriminator} | \uD83D\uDDD3 ${
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
    open fun onCommand(req: SlashCommandInteraction) {
        req.createImmediateResponder().setContent("Command not implemented!").setFlags(MessageFlag.EPHEMERAL).respond()
    }
}
