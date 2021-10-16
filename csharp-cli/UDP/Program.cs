using System;
using System.Net;
using System.Net.Sockets;
using System.Text;

namespace UDP
{
    class Program
    {
        static void Main(string[] args)
        {
            UDPcli();
        }

        public static void UDPcli()
        {
UdpClient udpClient = new UdpClient();

Byte[] sendBytes = Encoding.ASCII.GetBytes("HOWDY SEREGA!!!"); 
try{
    udpClient.Send(sendBytes, sendBytes.Length, "127.0.0.1", 8080);
}
catch ( Exception e ){
    Console.WriteLine(e.ToString());	
}
        }
    }
}
    
