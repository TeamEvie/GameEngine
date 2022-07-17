package utils

import java.io.File

class EnvUtils {
    companion object {
        /** Get an environment variable, or if it is not set, find it in the .env file. */
        fun getEnv(key: String): String {
            try {
                val value = System.getenv(key)

                if (value != null) {
                    return value
                }
            } catch (e: Exception) {}

            val file = File(".env")

            if (file.exists()) {
                val lines = file.readLines()

                for (line in lines) {
                    if (line.startsWith(key)) {
                        return line.substring(key.length + 1)
                    }
                }
            }
            throw IllegalArgumentException("Environment variable $key is not set.")
        }
    }
}
