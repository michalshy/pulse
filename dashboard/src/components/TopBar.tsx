
function NavBar({ isConnected }: { isConnected: boolean }) {
    return (
        <nav>
            <span>PULSE</span>
            <span style={{ color: isConnected ? '#00ff94' : #FF4444}}>
                {isConnected ? 'CONNECTED' : 'DISCONNECTED'}
            </span>
        </nav>
    )
}