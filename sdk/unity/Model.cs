string MetricToJson(MetricEntry m)
{
    return $"{{\"name\":\"{m.name}\",\"recorded_at\":\"{m.recorded_at}\",\"value_type\":\"{m.value_type}\",\"value\":{m.value}}}";
}

[Serializable]
public class HandshakeMessgae
{
    public string type = "handshake";
    public string project_key;
}


[Serializable]
public class MetricEntry
{
    public string name;
    public string recorded_at;
    public string value_type;  // "float", "int", "vec2", "vec3", "string", "json"
    public string value;
}

[Serializable]
public class EventEntry
{
    public string name;
    public string recorded_at;
    public string metadata;
}

[Serializable]
public class FlushMessage
{
    public string type = "flush";
    public MetricEntry[] metrics;
    public EventEntry[] events;
}