import java.net.ServerSocket;
import java.net.Socket;
import java.io.InputStream;
import java.io.OutputStream;
import java.io.InputStreamReader;
import java.io.BufferedReader;
import java.text.*;
import java.util.Calendar;
import java.util.TimeZone;

public class HttpServer {

    private static SimpleDateFormat ft = new SimpleDateFormat("E, dd MMM yyyy hh:mm:ss z");

    public static void main(String[] args) throws Throwable {
        ft.setTimeZone(TimeZone.getTimeZone("GMT"));
        ServerSocket ss = new ServerSocket(7000);
        while (true) {
            Socket s = ss.accept();
            new Thread(new SocketProcessor(s)).start();
        }
    }

    private static class SocketProcessor implements Runnable {

        private Socket s;
        private InputStream is;
        private OutputStream os;

        private SocketProcessor(Socket s) throws Throwable {
            this.s = s;
            this.is = s.getInputStream();
            this.os = s.getOutputStream();
        }

        public void run() {
            try {
                readInputHeaders();
                writeResponse("Hello, World!\n");
            } catch (Throwable t) {
                /*do nothing*/
            } finally {
                try {
                    s.close();
                } catch (Throwable t) {
                    /*do nothing*/
                }
            }
        }

        private void writeResponse(String s) throws Throwable {
            String response = "HTTP/1.1 200 OK\r\n" +
                    "Date: " + ft.format(Calendar.getInstance().getTime()) + "\r\n" +
                    "Content-Length: " + s.length() + "\r\n" +
                    "Content-Type: text/plain; charset=utf-8\r\n\r\n";
            String result = response + s;
            os.write(result.getBytes());
            os.flush();
        }

        private void readInputHeaders() throws Throwable {
            BufferedReader br = new BufferedReader(new InputStreamReader(is));
            while(true) {
                String s = br.readLine();
                if(s == null || s.trim().length() == 0) {
                    break;
                }
            }
        }
    }
}
